package main

import (
	"os"

	common "github.com/noironetworks/cilium-net/common"
	"github.com/noironetworks/cilium-net/plugins/cilium-docker/driver"

	"github.com/codegangsta/cli"
	l "github.com/op/go-logging"
)

const (
	// PluginPath is the docker plugins directory where docker plugin is present.
	pluginPath = "/run/docker/plugins/"
	// driverSock is the cilium socket for the communication between docker and cilium.
	driverSock = pluginPath + "cilium.sock"
)

var log = l.MustGetLogger("cilium-net-docker-plugin")

func main() {
	app := cli.NewApp()
	app.Name = "cilium-net"
	app.Usage = "Cilium Networking Docker Plugin"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug messages",
		},
	}
	app.Before = initEnv
	app.Action = run
	app.Run(os.Args)
}

func initEnv(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		common.SetupLOG(log, "DEBUG")
	} else {
		common.SetupLOG(log, "INFO")
	}

	if err := os.MkdirAll(pluginPath, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("Could not create net plugin path directory: %s", err)
	}

	if _, err := os.Stat(driverSock); err == nil {
		log.Debugf("socket file %s already exists, unlinking the old file handle.", driverSock)
		os.RemoveAll(driverSock)
	}

	log.Debugf("The plugin absolute path and handle is %s", driverSock)

	return nil
}

func run(ctx *cli.Context) {
	d, err := driver.NewDriver(ctx)
	if err != nil {
		log.Fatalf("Unable to create cilium-net driver: %s", err)
	}

	log.Info("Cilium networking Docker plugin ready")

	if err := d.Listen(driverSock); err != nil {
		log.Fatal(err)
	}
}
