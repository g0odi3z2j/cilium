package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/noironetworks/cilium-net/common/ipam"
)

// AllocateIP sends a POST request to allocate a new IP for the given options to the
// daemon. Returns an IPAMConfig if the daemon returns a http.StatusCreated, which means
// the allocation was successfully made.
func (cli Client) AllocateIP(ipamType ipam.IPAMType, options ipam.IPAMReq) (*ipam.IPAMRep, error) {

	serverResp, err := cli.R().SetBody(options).Post("/allocator/ipam-allocate/" + string(ipamType))
	if err != nil {
		return nil, fmt.Errorf("error while connecting to daemon: %s", err)
	}

	if serverResp.StatusCode() != http.StatusCreated &&
		serverResp.StatusCode() != http.StatusNoContent {
		return nil, processErrorBody(serverResp.Body(), nil)
	}

	if serverResp.StatusCode() == http.StatusNoContent {
		return nil, nil
	}

	var newIPAMConfig ipam.IPAMRep
	if err := json.Unmarshal(serverResp.Body(), &newIPAMConfig); err != nil {
		return nil, err
	}

	return &newIPAMConfig, nil
}

// ReleaseIP sends a POST request to release the IP of the given options.
func (cli Client) ReleaseIP(ipamType ipam.IPAMType, options ipam.IPAMReq) error {

	serverResp, err := cli.R().SetBody(options).Post("/allocator/ipam-release/" + string(ipamType))
	if err != nil {
		return fmt.Errorf("error while connecting to daemon: %s", err)
	}

	if serverResp.StatusCode() != http.StatusNoContent {
		return processErrorBody(serverResp.Body(), nil)
	}

	return nil
}

// GetIPAMConf sends a POST request to retrieve the IPAM configuration for the given
// ipamType.
func (cli Client) GetIPAMConf(ipamType ipam.IPAMType, options ipam.IPAMReq) (*ipam.IPAMConfigRep, error) {

	serverResp, err := cli.R().SetBody(options).Post("/allocator/ipam-configuration/" + string(ipamType))
	if err != nil {
		return nil, fmt.Errorf("error while connecting to daemon: %s", err)
	}

	if serverResp.StatusCode() != http.StatusOK {
		return nil, processErrorBody(serverResp.Body(), nil)
	}

	var newIPAMRep ipam.IPAMConfigRep
	if err := json.Unmarshal(serverResp.Body(), &newIPAMRep); err != nil {
		return nil, err
	}

	return &newIPAMRep, nil
}
