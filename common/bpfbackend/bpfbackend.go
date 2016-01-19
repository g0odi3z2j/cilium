package bpfbackend

import (
	"os/exec"
	"strconv"

	common "github.com/noironetworks/cilium-net/common"
	ciliumtype "github.com/noironetworks/cilium-net/common/types"

	log "github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

func EndpointJoin(ep *ciliumtype.Endpoint) error {
	id := strconv.Itoa(common.EndpointID(ep.LxcIP))
	args := []string{id, ep.Ifname, ep.LxcMAC.String(), ep.LxcIP.String()}
	out, err := exec.Command("../common/bpf/join_ep.sh", args...).CombinedOutput()
	if err != nil {
		log.Warnf("Command execution failed: %s", err)
		log.Warnf("Command output:\n%s", out)
		return err
	}
	log.Infof("Command successful:\n%s", out)

	return nil
}

func EndpointLeave(ep *ciliumtype.Endpoint) error {
	id := strconv.Itoa(common.EndpointID(ep.LxcIP))
	args := []string{id, ep.Ifname, ep.LxcMAC.String(), ep.LxcIP.String()}
	out, err := exec.Command("../common/bpf/leave_ep.sh", args...).CombinedOutput()
	if err != nil {
		log.Warnf("Command execution failed: %s", err)
		log.Warnf("Command output:\n%s", out)
		return err
	}
	log.Infof("Command successful:\n%s", out)

	return nil
}
