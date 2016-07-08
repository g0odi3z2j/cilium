package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/noironetworks/cilium-net/bpf/policymap"
	"github.com/noironetworks/cilium-net/common"
)

// EPPortMap is the port mapping representation for a particular endpoint.
type EPPortMap struct {
	From  uint16 `json:"from"`
	To    uint16 `json:"to"`
	Proto uint8  `json:"proto"`
}

const (
	OptionNAT46               = "NAT46"
	OptionPolicy              = "Policy"
	OptionDropNotify          = "DropNotification"
	OptionConntrack           = "Conntrack"
	OptionConntrackAccounting = "ConntrackAccounting"
	OptionDebug               = "Debug"
	OptionAllowToHost         = "AllowToHost"
	OptionAllowToWorld        = "AllowToWorld"
	OptionLearnTraffic        = "LearnTraffic"
)

var (
	OptionSpecNAT46 = Option{
		Define:      "ENABLE_NAT46",
		Description: "Enable automatic NAT46 translation",
	}

	OptionSpecPolicy = Option{
		Define:      "POLICY_ENFORCEMENT",
		Description: "Enable policy enforcement",
	}

	OptionSpecDropNotify = Option{
		Define:      "DROP_NOTIFY",
		Description: "Enable drop notifications",
	}

	OptionSpecConntrack = Option{
		Define:      "CONNTRACK",
		Description: "Enable stateful connection tracking",
	}

	OptionSpecConntrackAccounting = Option{
		Define:      "CONNTRACK_ACCOUNTING",
		Description: "Enable per flow (conntrack) statistics",
	}

	OptionSpecDebug = Option{
		Define:      "DEBUG",
		Description: "Enable debugging trace statements",
	}

	OptionSpecAllowToHost = Option{
		Define:      "ALLOW_TO_HOST",
		Immutable:   true,
		Description: "Allow all traffic to local host",
	}

	OptionSpecAllowToWorld = Option{
		Define:      "ALLOW_TO_WORLD",
		Immutable:   true,
		Description: "Allow all traffic to outside world",
	}

	OptionSpecLearnTraffic = Option{
		Define:      "LEARN_TRAFFIC",
		Description: "Learn and add labels to the list of allowed labels",
	}

	EndpointOptionLibrary = OptionLibrary{
		OptionNAT46:               &OptionSpecNAT46,
		OptionPolicy:              &OptionSpecPolicy,
		OptionDropNotify:          &OptionSpecDropNotify,
		OptionConntrack:           &OptionSpecConntrack,
		OptionConntrackAccounting: &OptionSpecConntrackAccounting,
		OptionDebug:               &OptionSpecDebug,
		OptionAllowToHost:         &OptionSpecAllowToHost,
		OptionAllowToWorld:        &OptionSpecAllowToWorld,
		OptionLearnTraffic:        &OptionSpecLearnTraffic,
	}
)

// Endpoint contains all the details for a particular LXC and the host interface to where
// is connected to.
type Endpoint struct {
	ID               uint16               `json:"id"`                 // Endpoint ID.
	DockerID         string               `json:"docker-id"`          // Docker ID.
	DockerNetworkID  string               `json:"docker-network-id"`  // Docker network ID.
	DockerEndpointID string               `json:"docker-endpoint-id"` // Docker endpoint ID.
	IfName           string               `json:"interface-name"`     // Container's interface name.
	LXCMAC           MAC                  `json:"lxc-mac"`            // Container MAC address.
	LXCIP            net.IP               `json:"lxc-ip"`             // Container IPv6 address.
	IfIndex          int                  `json:"interface-index"`    // Host's interface index.
	NodeMAC          MAC                  `json:"node-mac"`           // Node MAC address.
	NodeIP           net.IP               `json:"node-ip"`            // Node IPv6 address.
	SecLabel         *SecCtxLabel         `json:"security-label"`     // Security Label  set to this endpoint.
	PortMap          []EPPortMap          `json:"port-mapping"`       // Port mapping used for this endpoint.
	Consumable       *Consumable          `json:"consumable"`
	PolicyMap        *policymap.PolicyMap `json:"-"`
	Opts             *BoolOptions         `json:"options"` // Endpoint bpf options.
}

// SetID sets the endpoint's host local unique ID.
func (e *Endpoint) SetID() {
	e.ID = common.EndpointAddr2ID(e.LXCIP)
}

func (e *Endpoint) SetSecLabel(labels *SecCtxLabel) {
	e.SecLabel = labels
	e.Consumable = GetConsumable(labels.ID, labels)
}

func (e *Endpoint) Allows(id uint32) bool {
	if e.Consumable != nil {
		return e.Consumable.Allows(id)
	} else {
		return false
	}
}

// IPv4Address returns the TODO: what does this do?
func (e *Endpoint) IPv4Address(v4Range *net.IPNet) *net.IP {
	ip := common.DupIP(v4Range.IP)

	id := e.ID
	ip[2] = byte(id >> 8)
	ip[3] = byte(id & 0xff)

	return &ip
}

// String returns endpoint on a JSON format.
func (e Endpoint) String() string {
	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func OptionChanged(key string, value bool, data interface{}) {
	e := data.(*Endpoint)
	switch key {
	case OptionConntrack:
		e.InvalidatePolicy()
	}
}

func (e *Endpoint) ApplyOpts(opts OptionMap) bool {
	if val, ok := opts[OptionLearnTraffic]; ok && val {
		opts[OptionDropNotify] = true
	}

	return e.Opts.Apply(opts, OptionChanged, e) > 0
}

func (ep *Endpoint) SetDefaultOpts(opts *BoolOptions) {
	restore := false
	if ep.Opts == nil {
		restore = true
		ep.Opts = NewBoolOptions(&EndpointOptionLibrary)
	}
	if ep.Opts.Library == nil {
		ep.Opts.Library = &EndpointOptionLibrary
	}

	if restore {
		if opts != nil {
			ep.Opts.InheritDefault(opts, OptionConntrack)
			ep.Opts.InheritDefault(opts, OptionConntrackAccounting)
			ep.Opts.InheritDefault(opts, OptionPolicy)
			ep.Opts.InheritDefault(opts, OptionDebug)
			ep.Opts.InheritDefault(opts, OptionDropNotify)
			ep.Opts.InheritDefault(opts, OptionNAT46)
		}
		ep.Opts.SetIfUnset(OptionLearnTraffic, false)
	}
}

type orderEndpoint func(e1, e2 *Endpoint) bool

// OrderEndpointAsc orders the slice of Endpoint in ascending ID order.
func OrderEndpointAsc(eps []Endpoint) {
	ascPriority := func(e1, e2 *Endpoint) bool {
		return e1.ID < e2.ID
	}
	orderEndpoint(ascPriority).sort(eps)
}

func (by orderEndpoint) sort(eps []Endpoint) {
	dS := &epSorter{
		eps: eps,
		by:  by,
	}
	sort.Sort(dS)
}

type epSorter struct {
	eps []Endpoint
	by  func(e1, e2 *Endpoint) bool
}

func (epS *epSorter) Len() int {
	return len(epS.eps)
}

func (epS *epSorter) Swap(i, j int) {
	epS.eps[i], epS.eps[j] = epS.eps[j], epS.eps[i]
}

func (epS *epSorter) Less(i, j int) bool {
	return epS.by(&epS.eps[i], &epS.eps[j])
}

// Base64 returns the endpoint in a base64 format.
func (ep Endpoint) Base64() (string, error) {
	jsonBytes, err := json.Marshal(ep)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

// ParseBase64ToEndpoint parses the endpoint stored in the given base64 string.
func ParseBase64ToEndpoint(str string, ep *Endpoint) error {
	jsonBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, ep)
}

// FilterEPDir returns a list of directories' names that possible belong to an endpoint.
func FilterEPDir(dirFiles []os.FileInfo) []string {
	eptsID := []string{}
	for _, file := range dirFiles {
		if file.IsDir() {
			if _, err := strconv.ParseUint(file.Name(), 10, 16); err == nil {
				eptsID = append(eptsID, file.Name())
			}
		}
	}
	return eptsID
}

// ParseEndpoint parses the given strEp which is in the form of:
// common.CiliumCHeaderPrefix + common.Version + ":" + endpointBase64
func ParseEndpoint(strEp string) (*Endpoint, error) {
	// TODO: Provide a better mechanism to update from old version once we bump
	// TODO: cilium version.
	strEpSlice := strings.Split(strEp, ":")
	if len(strEpSlice) != 2 {
		return nil, fmt.Errorf("invalid format %q. Should contain a single ':'", strEp)
	}
	var ep Endpoint
	if err := ParseBase64ToEndpoint(strEpSlice[1], &ep); err != nil {
		return nil, fmt.Errorf("failed to parse base64toendpoint: %s", err)
	}
	return &ep, nil
}

// IsLibnetwork returns true if the endpoint was created by Libnetwork, false otherwise.
func (e *Endpoint) IsLibnetwork() bool {
	return e.DockerNetworkID != ""
}

// IsCNI returns true if the endpoint was created by CNI, false otherwise.
func (e *Endpoint) IsCNI() bool {
	return e.DockerNetworkID == ""
}

func (e *Endpoint) PolicyMapPath() string {
	return common.PolicyMapPath + strconv.Itoa(int(e.ID))
}

func (e *Endpoint) InvalidatePolicy() {
	if e.Consumable != nil {
		// Resetting to 0 will trigger a regeneration on the next update
		log.Debugf("Invalidated policy for endpoint %s", e.ID)
		e.Consumable.Iteration = 0
	}
}
