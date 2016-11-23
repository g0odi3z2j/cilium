//
// Copyright 2016 Authors of Cilium
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
//
package daemon

import (
	"fmt"

	"github.com/cilium/cilium/pkg/endpoint"

	ctx "golang.org/x/net/context"
	k8sTypes "k8s.io/client-go/1.5/pkg/api/v1"
)

func (d *Daemon) getK8sStatus() endpoint.Status {
	var k8sStatus endpoint.Status
	if d.conf.IsK8sEnabled() {
		if v, err := d.k8sClient.ComponentStatuses().Get("controller-manager"); err != nil {
			k8sStatus = endpoint.Status{Code: endpoint.Failure, Msg: err.Error()}
		} else if len(v.Conditions) == 0 {
			k8sStatus = endpoint.Status{Code: endpoint.Warning, Msg: "Unable to retrieve controller-manager's kubernetes status"}
		} else {
			if v.Conditions[0].Status == k8sTypes.ConditionTrue {
				k8sStatus = endpoint.NewStatusOK(string(v.Conditions[0].Type))
			} else {
				k8sStatus = endpoint.Status{Code: endpoint.Failure, Msg: v.Conditions[0].Message}
			}
		}
	} else {
		k8sStatus = endpoint.Status{Code: endpoint.Disabled}
	}
	return k8sStatus
}

func (d *Daemon) GlobalStatus() (*endpoint.StatusResponse, error) {
	sr := endpoint.StatusResponse{}
	if info, err := d.kvClient.Status(); err != nil {
		sr.KVStore = endpoint.Status{Code: endpoint.Failure, Msg: fmt.Sprintf("Err: %s - %s", err, info)}
	} else {
		sr.KVStore = endpoint.NewStatusOK(info)
	}

	if _, err := d.dockerClient.Info(ctx.Background()); err != nil {
		sr.Docker = endpoint.Status{Code: endpoint.Failure, Msg: err.Error()}
	} else {
		sr.Docker = endpoint.NewStatusOK("")
	}

	sr.Kubernetes = d.getK8sStatus()

	if sr.KVStore.Code != endpoint.OK {
		sr.Cilium = endpoint.Status{Code: sr.KVStore.Code, Msg: "KVStore service is not ready!"}
	} else if sr.Docker.Code != endpoint.OK {
		sr.Cilium = endpoint.Status{Code: sr.Docker.Code, Msg: "Docker service is not ready!"}
	} else if d.conf.IsK8sEnabled() && sr.Kubernetes.Code != endpoint.OK {
		sr.Cilium = endpoint.Status{Code: sr.Kubernetes.Code, Msg: "Kubernetes service is not ready!"}
	}

	// TODO Create a logstash status in its runnable function
	//Logstash   Status `json:"logstash"`

	sr.IPAMStatus = d.DumpIPAM()

	return &sr, nil
}
