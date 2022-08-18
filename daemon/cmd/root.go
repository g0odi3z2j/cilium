// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"fmt"
	"os"
	"time"

	gops "github.com/google/gops/agent"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/cilium/cilium/pkg/hive"
	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/version"
)

var (
	RootCmd = &cobra.Command{
		Use:   "cilium-agent",
		Short: "Run the cilium agent",
		Run:   runApp,
	}

	cmdrefCmd = &cobra.Command{
		Use:   "cmdref [output directory]",
		Short: "Generate command reference for cilium-agent to given output directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			genMarkdown(RootCmd, args[0])
		},
	}

	dotGraphCmd = &cobra.Command{
		Use:   "dot-graph",
		Short: "Output the internal dependencies of cilium-agent in graphviz dot format",
		Run:   func(cmd *cobra.Command, args []string) { agentHive.PrintDotGraph() },
	}

	objectsCmd = &cobra.Command{
		Use:   "objects",
		Short: "Print the objects, constructors and lifecycle hooks",
		Run:   func(cmd *cobra.Command, args []string) { agentHive.PrintObjects() },
	}

	agentHive *hive.Hive
)

func init() {
	cobra.OnInitialize(option.InitConfig(RootCmd, "cilium-agent", "cilium", Vp))
	setupSleepBeforeFatal()
	registerBootstrapMetrics()
	initializeFlags()

	agentHive = hive.New(
		Vp, RootCmd.Flags(),

		hive.NewCell("daemon", fx.Invoke(registerDaemonHooks)),
	)
}

func runApp(cmd *cobra.Command, args []string) {
	bootstrapStats.overall.Start()

	if v, _ := cmd.Flags().GetBool("version"); v {
		fmt.Printf("%s %s\n", cmd.Name(), version.Version)
		os.Exit(0)
	}

	// Open socket for using gops to get stacktraces of the agent.
	addr := fmt.Sprintf("127.0.0.1:%d", Vp.GetInt(option.GopsPort))
	addrField := logrus.Fields{"address": addr}
	if err := gops.Listen(gops.Options{
		Addr:                   addr,
		ReuseSocketAddrAndPort: true,
	}); err != nil {
		log.WithError(err).WithFields(addrField).Fatal("Cannot start gops server")
	}
	log.WithFields(addrField).Info("Started gops server")

	agentHive.Run()
}

func Execute() error {
	RootCmd.AddCommand(
		cmdrefCmd,
		dotGraphCmd,
		objectsCmd,
	)

	return RootCmd.Execute()
}

func setupSleepBeforeFatal() {
	RootCmd.SetFlagErrorFunc(
		func(_ *cobra.Command, e error) error {
			time.Sleep(fatalSleep)
			return e
		})
	logrus.RegisterExitHandler(func() {
		time.Sleep(fatalSleep)
	})
}
