// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package loader

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/cilium/cilium/pkg/command/exec"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/option"

	"github.com/sirupsen/logrus"
)

// OutputType determines the type to be generated by the compilation steps.
type OutputType string

const (
	outputObject   = OutputType("obj")
	outputAssembly = OutputType("asm")
	outputSource   = OutputType("c")
)

// progInfo describes a program to be compiled with the expected output format
type progInfo struct {
	// Source is the program source (base) filename to be compiled
	Source string
	// Output is the expected (base) filename produced from the source
	Output string
	// OutputType to be created by LLVM
	OutputType OutputType
}

// directoryInfo includes relevant directories for compilation and linking
type directoryInfo struct {
	// Library contains the library code to be used for compilation
	Library string
	// Runtime contains headers for compilation
	Runtime string
	// State contains node, lxc, and features headers for templatization
	State string
	// Output is the directory where the files will be stored
	Output string
}

var (
	compiler       = "clang"
	linker         = "llc"
	standardCFlags = []string{"-O2", "-target", "bpf",
		fmt.Sprintf("-D__NR_CPUS__=%d", runtime.NumCPU()),
		"-Wno-address-of-packed-member", "-Wno-unknown-warning-option"}
	standardLDFlags = []string{"-march=bpf", "-mcpu=probe"}

	endpointPrefix   = "bpf_lxc"
	endpointProg     = fmt.Sprintf("%s.%s", endpointPrefix, outputSource)
	endpointObj      = fmt.Sprintf("%s.o", endpointPrefix)
	endpointObjDebug = fmt.Sprintf("%s.dbg.o", endpointPrefix)
	endpointAsm      = fmt.Sprintf("%s.%s", endpointPrefix, outputAssembly)

	// testIncludes allows the unit tests to inject additional include
	// paths into the compile command at test time. It is usually nil.
	testIncludes []string

	debugProgs = []*progInfo{
		{
			Source:     endpointProg,
			Output:     endpointObjDebug,
			OutputType: outputObject,
		},
		{
			Source:     endpointProg,
			Output:     endpointAsm,
			OutputType: outputAssembly,
		},
		{
			Source:     endpointProg,
			Output:     endpointProg,
			OutputType: outputSource,
		},
	}
	datapathProg = &progInfo{
		Source:     endpointProg,
		Output:     endpointObj,
		OutputType: outputObject,
	}
)

// progLDFlags determines the loader flags for the specified prog and paths.
func progLDFlags(prog *progInfo, dir *directoryInfo) []string {
	return []string{
		fmt.Sprintf("-filetype=%s", prog.OutputType),
		"-o", path.Join(dir.Output, prog.Output),
	}
}

// prepareCmdPipes attaches pipes to the stdout and stderr of the specified
// command, and returns the stdout, stderr, and any error that may have
// occurred while creating the pipes.
func prepareCmdPipes(cmd *exec.Cmd) (io.ReadCloser, io.ReadCloser, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get stdout pipe: %s", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		stdout.Close()
		return nil, nil, fmt.Errorf("Failed to get stderr pipe: %s", err)
	}

	return stdout, stderr, nil
}

func pidFromProcess(proc *os.Process) string {
	result := "not-started"
	if proc != nil {
		result = fmt.Sprintf("%d", proc.Pid)
	}
	return result
}

// compileAndLink links the specified program from the specified path to the
// intermediate representation, to the output specified in the prog's info.
func compileAndLink(ctx context.Context, prog *progInfo, dir *directoryInfo, debug bool, compileArgs ...string) error {
	compileCmd, cancelCompile := exec.WithCancel(ctx, compiler, compileArgs...)
	defer cancelCompile()
	compilerStdout, compilerStderr, err := prepareCmdPipes(compileCmd)
	if err != nil {
		return err
	}

	linkArgs := make([]string, 0, 8)
	if debug {
		linkArgs = append(linkArgs, "-mattr=dwarfris")
	}
	linkArgs = append(linkArgs, standardLDFlags...)
	linkArgs = append(linkArgs, progLDFlags(prog, dir)...)

	linkCmd := exec.CommandContext(ctx, linker, linkArgs...)
	linkCmd.Stdin = compilerStdout
	if err := compileCmd.Start(); err != nil {
		return fmt.Errorf("Failed to start command %s: %s", compileCmd.Args, err)
	}

	var compileOut []byte
	/* Ignoring the output here because pkg/command/exec will log it. */
	_, err = linkCmd.CombinedOutput(log, true)
	if err == nil {
		compileOut, _ = ioutil.ReadAll(compilerStderr)
		err = compileCmd.Wait()
	} else {
		cancelCompile()
	}
	if err != nil {
		err = fmt.Errorf("Failed to compile %s: %s", prog.Output, err)
		log.WithFields(logrus.Fields{
			"compiler-pid": pidFromProcess(compileCmd.Process),
			"linker-pid":   pidFromProcess(linkCmd.Process),
		}).Error(err)
		if compileOut != nil {
			scopedLog := log.Warn
			if debug {
				scopedLog = log.Debug
			}
			scanner := bufio.NewScanner(bytes.NewReader(compileOut))
			for scanner.Scan() {
				scopedLog(scanner.Text())
			}
		}
	}

	return err
}

// progLDFlags determines the compiler flags for the specified prog and paths.
func progCFlags(prog *progInfo, dir *directoryInfo) []string {
	var output string

	if prog.OutputType == outputSource {
		output = path.Join(dir.Output, prog.Output)
	} else {
		output = "-" // stdout
	}

	return append(testIncludes,
		fmt.Sprintf("-I%s", path.Join(dir.Runtime, "globals")),
		fmt.Sprintf("-I%s", dir.State),
		fmt.Sprintf("-I%s", path.Join(dir.Library, "include")),
		"-c", path.Join(dir.Library, prog.Source),
		"-o", output,
	)
}

// compile and link a program.
func compile(ctx context.Context, prog *progInfo, dir *directoryInfo, debug bool) (err error) {
	args := make([]string, 0, 16)
	if prog.OutputType == outputSource {
		args = append(args, "-E") // Preprocessor
	} else {
		args = append(args, "-emit-llvm")
		if debug {
			args = append(args, "-g")
		}
	}
	args = append(args, standardCFlags...)
	args = append(args, progCFlags(prog, dir)...)

	// Compilation is split between two exec calls. First clang generates
	// LLVM bitcode and then later llc compiles it to byte-code.
	log.WithFields(logrus.Fields{
		"target": compiler,
		"args":   args,
	}).Debug("Launching compiler")
	if prog.OutputType == outputSource {
		compileCmd := exec.CommandContext(ctx, compiler, args...)
		_, err = compileCmd.CombinedOutput(log, debug)
	} else {
		switch prog.OutputType {
		case outputObject:
			err = compileAndLink(ctx, prog, dir, debug, args...)
		case outputAssembly:
			err = compileAndLink(ctx, prog, dir, false, args...)
		default:
			log.Fatalf("Unhandled progInfo.OutputType %s", prog.OutputType)
		}
	}

	return err
}

// compileDatapath invokes the compiler and linker to create all state files for
// the BPF datapath, with the primary target being the BPF ELF binary.
//
// If debug is enabled, create also the following output files:
// * Preprocessed C
// * Assembly
// * Object compiled with debug symbols
func compileDatapath(ctx context.Context, ep endpoint, dirs *directoryInfo, debug bool) error {
	// TODO: Consider logging kernel/clang versions here too
	epLog := ep.Logger(Subsystem)

	// Write out assembly and preprocessing files for debugging purposes
	if debug {
		for _, p := range debugProgs {
			if err := compile(ctx, p, dirs, debug); err != nil {
				scopedLog := epLog.WithFields(logrus.Fields{
					logfields.Params: logfields.Repr(p),
					logfields.Debug:  debug,
				})
				scopedLog.WithError(err).Debug("JoinEP: Failed to compile")
				return err
			}
		}
	}

	// Compile the new program
	if err := compile(ctx, datapathProg, dirs, debug); err != nil {
		scopedLog := epLog.WithFields(logrus.Fields{
			logfields.Params: logfields.Repr(datapathProg),
			logfields.Debug:  false,
		})
		scopedLog.WithError(err).Warn("JoinEP: Failed to compile")
		return err
	}

	return nil
}

// Compile compiles a BPF program generating an object file.
func Compile(ctx context.Context, src string, out string) error {
	debug := option.Config.BPFCompilationDebug
	prog := progInfo{
		Source:     src,
		Output:     out,
		OutputType: outputObject,
	}
	dirs := directoryInfo{
		Library: option.Config.BpfDir,
		Runtime: option.Config.StateDir,
		Output:  option.Config.StateDir,
		State:   option.Config.StateDir,
	}
	return compile(ctx, &prog, &dirs, debug)
}
