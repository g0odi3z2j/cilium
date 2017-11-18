// Copyright 2017 Authors of Cilium
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

package helpers

import (
	"github.com/sirupsen/logrus"
)

//CreateNewRuntimeHelper returns Docker and Cilium helpers for running the
//runtime tests on the provided VM target and using logger log .
func CreateNewRuntimeHelper(target string, log *logrus.Entry) (*Docker, *Cilium) {
	log.Infof("creating docker")
	docker := CreateDocker(target)
	log.Infof("creating cilium")
	cilium := CreateCilium(target, log)
	return docker, cilium
}
