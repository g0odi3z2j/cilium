//go:build !windows
// +build !windows

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"os"

	"github.com/prometheus/procfs"
)

func getProcessStart() (float64, error) {
	pid := os.Getpid()
	p, err := procfs.NewProc(pid)
	if err != nil {
		return 0, err
	}

	if stat, err := p.Stat(); err == nil {
		return stat.StartTime()
	}
	return 0, err
}
