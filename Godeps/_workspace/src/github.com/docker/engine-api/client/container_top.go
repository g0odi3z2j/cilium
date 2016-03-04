package client

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/docker/engine-api/types"
)

// ContainerTop shows process information from within a container.
func (cli *Client) ContainerTop(containerID string, arguments []string) (types.ContainerProcessList, error) {
	var response types.ContainerProcessList
	query := url.Values{}
	if len(arguments) > 0 {
		query.Set("ps_args", strings.Join(arguments, " "))
	}

	resp, err := cli.get("/containers/"+containerID+"/top", query, nil)
	if err != nil {
		return response, err
	}
	defer ensureReaderClosed(resp)

	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
