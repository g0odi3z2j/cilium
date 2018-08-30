// Copyright 2017-2018 Authors of Cilium
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
	"os"
	"os/exec"
	"path"
	"runtime"

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

// irPath determines a temporary path on the filesystem to store the compiled
// intermediate representation object for the program.
//
// TODO: Remove this and just pipe the 'clang' and 'llc' programs together.
func irPath(prog *progInfo, dir *directoryInfo) string {
	return fmt.Sprintf("%s/%s.bc", dir.Output, prog.Output)
}

// progLDFlags determines the loader flags for the specified prog and paths.
func progLDFlags(prog *progInfo, dir *directoryInfo) []string {
	return []string{
		fmt.Sprintf("-filetype=%s", prog.OutputType),
		irPath(prog, dir),
		"-o", path.Join(dir.Output, prog.Output),
	}
}

// linkProg links the specified program from the specified path to the
// intermediate representation, to the output specified in the prog's info.
func linkProg(ctx context.Context, prog *progInfo, dir *directoryInfo, debug bool) error {
	args := make([]string, 0, 8)
	if debug {
		args = append(args, "-mattr=dwarfris")
	}
	args = append(args, standardLDFlags...)
	args = append(args, progLDFlags(prog, dir)...)

	out, err := exec.CommandContext(ctx, linker, args...).CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		err = fmt.Errorf("Command execution failed: Timeout")
	}
	if err != nil {
		scanner := bufio.NewScanner(bytes.NewReader(out))
		return fmt.Errorf("Failed to link %s: %s: %q", prog.Output, err, scanner.Text())
	}

	return nil
}

// progLDFlags determines the compiler flags for the specified prog and paths.
func progCFlags(prog *progInfo, dir *directoryInfo) []string {
	var output string

	if prog.OutputType == outputSource {
		output = path.Join(dir.Output, prog.Output)
	} else {
		output = irPath(prog, dir)
	}

	return []string{
		fmt.Sprintf("-I%s", path.Join(dir.Runtime, "globals")),
		fmt.Sprintf("-I%s", dir.Output),
		fmt.Sprintf("-I%s", path.Join(dir.Library, "include")),
		"-c", path.Join(dir.Library, prog.Source),
		"-o", output,
	}
}

// compile and link a program.
func compile(ctx context.Context, prog *progInfo, dir *directoryInfo, debug bool) error {
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
	}).Debug("Launching CommandContext")
	out, err := exec.CommandContext(ctx, compiler, args...).CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		err = fmt.Errorf("Command execution failed: Timeout")
	}
	if err != nil {
		scanner := bufio.NewScanner(bytes.NewReader(out))
		return fmt.Errorf("Failed to compile %s: %s: %q", prog.Output, err, scanner.Text())
	}
	defer os.Remove(irPath(prog, dir))

	switch prog.OutputType {
	case outputObject:
		return linkProg(ctx, prog, dir, debug)
	case outputAssembly:
		return linkProg(ctx, prog, dir, false)
	default:
		break
	}

	return nil
}
