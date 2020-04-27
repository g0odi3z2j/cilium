// Copyright 2020 Authors of Cilium
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

package serve

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/cilium/cilium/pkg/hubble/relay"
	"github.com/cilium/cilium/pkg/hubble/relay/relayoption"
	"golang.org/x/sys/unix"

	"github.com/spf13/cobra"
)

var flag struct {
	debug         bool
	dialTimeout   time.Duration
	listenAddress string
	retryTimeout  time.Duration
}

// New creates a new serve command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the gRPC proxy server",
		Long:  `Run the gRPC proxy server.`,
		RunE: func(_ *cobra.Command, _ []string) error {
			return runServe()
		},
	}
	cmd.Flags().BoolVarP(&flag.debug, "debug", "D", false, "Run in debug mode")
	cmd.Flags().DurationVar(&flag.dialTimeout, "dial-timeout", relayoption.Default.DialTimeout, "Dial timeout when connecting to hubble peers")
	cmd.Flags().DurationVar(&flag.retryTimeout, "retry-timeout", relayoption.Default.RetryTimeout, "Time to wait before attempting to reconnect to a hubble peer when the connection is lost")
	cmd.Flags().StringVar(&flag.listenAddress, "listen-address", relayoption.Default.ListenAddress, "Address on which to listen")
	return cmd
}

func runServe() error {
	opts := []relayoption.Option{
		relayoption.WithDialTimeout(flag.dialTimeout),
		relayoption.WithListenAddress(flag.listenAddress),
		relayoption.WithRetryTimeout(flag.retryTimeout),
	}
	if flag.debug {
		opts = append(opts, relayoption.WithDebug())
	}
	srv, err := relay.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("cannot create hubble-relay server: %v", err)
	}
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, unix.SIGINT, unix.SIGTERM)
		<-sigs
		srv.Stop()
	}()
	return srv.Serve()
}
