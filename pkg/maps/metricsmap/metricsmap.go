// Copyright 2016-2020 Authors of Cilium
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

package metricsmap

import (
	"context"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/cilium/cilium/pkg/bpf"
	"github.com/cilium/cilium/pkg/common"
	"github.com/cilium/cilium/pkg/logging"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/metrics"
	monitorAPI "github.com/cilium/cilium/pkg/monitor/api"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Metrics is the bpf metrics map
	Metrics      *bpf.Map
	log          = logging.DefaultLogger.WithField(logfields.LogSubsys, "map-metrics")
	possibleCpus int
)

const (
	// MapName for metrics map.
	MapName = "cilium_metrics"
	// MaxEntries is the maximum number of keys that can be present in the
	// Metrics Map.
	//
	// Currently max. 2 bits of the Key.Dir member are used (unknown,
	// ingress or egress). Thus we can reduce from the theoretical max. size
	// of 2**16 (2 uint8) to 2**10 (1 uint8 + 2 bits).
	MaxEntries = 1024
	// dirIngress and dirEgress values should match with
	// METRIC_INGRESS, METRIC_EGRESS and METRIC_SERVICE
	// in bpf/lib/common.h
	dirEgress  = 0
	dirIngress = 1
	dirService = 2
)

// direction is the metrics direction i.e ingress (to an endpoint),
// egress (from an endpoint) or service (NodePort service being accessed from
// outside or a ClusterIP service being accessed from inside the cluster).
// If it's none of the above, we return UNKNOWN direction.
var direction = map[uint8]string{
	dirEgress:  "EGRESS",
	dirIngress: "INGRESS",
	dirService: "SERVICE",
}

type pad3uint16 [3]uint16

// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *pad3uint16) DeepCopyInto(out *pad3uint16) {
	copy(out[:], in[:])
	return
}

// Key must be in sync with struct metrics_key in <bpf/lib/common.h>
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=github.com/cilium/cilium/pkg/bpf.MapKey
type Key struct {
	Reason   uint8      `align:"reason"`
	Dir      uint8      `align:"dir"`
	Reserved pad3uint16 `align:"reserved"`
}

// Value must be in sync with struct metrics_value in <bpf/lib/common.h>
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=github.com/cilium/cilium/pkg/bpf.MapValue
type Value struct {
	Count uint64 `align:"count"`
	Bytes uint64 `align:"bytes"`
}

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=github.com/cilium/cilium/pkg/bpf.MapValue
// Values is a slice of Values
type Values []Value

// Record designates a map entry (key + value).
// This type is used for mock maps.
type Record struct {
	Key   Key
	Value Value
}

// MetricsMap interface represents a metrics map, and can be reused to implement
// mock maps for unit tests.
type MetricsMap interface {
	DumpWithCallback(bpf.DumpCallback) error
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (vs *Values) DeepCopyMapValue() bpf.MapValue {
	if c := vs.DeepCopy(); c != nil {
		return &c
	}
	return nil
}

// String converts the value into a human readable string format
func (vs Values) String() string {
	sumCount, sumBytes := uint64(0), uint64(0)
	for _, v := range vs {
		sumCount += v.Count
		sumBytes += v.Bytes
	}
	return fmt.Sprintf("count:%d bytes:%d", sumCount, sumBytes)
}

// GetValuePtr returns the unsafe pointer to the BPF value.
func (vs *Values) GetValuePtr() unsafe.Pointer {
	return unsafe.Pointer(vs)
}

// String converts the key into a human readable string format
func (k *Key) String() string {
	return fmt.Sprintf("reason:%d dir:%d", k.Reason, k.Dir)
}

// MetricDirection gets the direction in human readable string format
func MetricDirection(dir uint8) string {
	if desc, ok := direction[dir]; ok {
		return desc
	}
	return "UNKNOWN"
}

// Direction gets the direction in human readable string format
func (k *Key) Direction() string {
	return MetricDirection(k.Dir)
}

// DropForwardReason gets the forwarded/dropped reason in human readable string format
func (k *Key) DropForwardReason() string {
	return monitorAPI.DropReason(k.Reason)
}

// GetKeyPtr returns the unsafe pointer to the BPF key
func (k *Key) GetKeyPtr() unsafe.Pointer { return unsafe.Pointer(k) }

// String converts the value into a human readable string format
func (v *Value) String() string {
	return fmt.Sprintf("count:%d bytes:%d", v.Count, v.Bytes)
}

// RequestCount returns the drop/forward count in a human readable string format
func (v *Value) RequestCount() string {
	return strconv.FormatUint(v.Count, 10)
}

// RequestBytes returns drop/forward bytes in a human readable string format
func (v *Value) RequestBytes() string {
	return strconv.FormatUint(v.Bytes, 10)
}

// IsDrop checks if the reason is drop or not.
func (k *Key) IsDrop() bool {
	return k.Reason == monitorAPI.DropInvalid || k.Reason >= monitorAPI.DropMin
}

// CountFloat converts the request count to float
func (v *Value) CountFloat() float64 {
	return float64(v.Count)
}

// bytesFloat converts the bytes count to float
func (v *Value) bytesFloat() float64 {
	return float64(v.Bytes)
}

// NewValue returns a new empty instance of the structure representing the BPF
// map value
func (k *Key) NewValue() bpf.MapValue { return &Value{} }

// GetValuePtr returns the unsafe pointer to the BPF value.
func (v *Value) GetValuePtr() unsafe.Pointer {
	return unsafe.Pointer(v)
}

func updateMetric(getCounter func() (prometheus.Counter, error), newValue float64) {
	counter, err := getCounter()
	if err != nil {
		log.WithError(err).Warn("Failed to update prometheus metrics")
		return
	}

	oldValue := metrics.GetCounterValue(counter)
	if newValue > oldValue {
		counter.Add(newValue - oldValue)
	}
}

// updatePrometheusMetrics checks the metricsmap key value pair
// and determines which prometheus metrics along with respective labels
// need to be updated.
func updatePrometheusMetrics(key *Key, val *Value) {
	updateMetric(func() (prometheus.Counter, error) {
		if key.IsDrop() {
			return metrics.DropCount.GetMetricWithLabelValues(key.DropForwardReason(), key.Direction())
		}
		return metrics.ForwardCount.GetMetricWithLabelValues(key.Direction())
	}, val.CountFloat())

	updateMetric(func() (prometheus.Counter, error) {
		if key.IsDrop() {
			return metrics.DropBytes.GetMetricWithLabelValues(key.DropForwardReason(), key.Direction())
		}
		return metrics.ForwardBytes.GetMetricWithLabelValues(key.Direction())
	}, val.bytesFloat())
}

// SyncMetricsMap is called periodically to sync off the metrics map by
// aggregating it into drops (by drop reason and direction) and
// forwards (by direction) with the prometheus server.
func SyncMetricsMap(ctx context.Context) error {
	entry := make([]Value, possibleCpus)
	file := bpf.MapPath(MapName)

	var err error
	metricsMap := bpf.GetMap(file)
	if metricsMap == nil {
		// Open the map and leave it open, since SyncMetricsMap is called
		// periodically and it makes sense to use an already opened map rather
		// than opening the map again and again.
		// This also prevents the constant registration and unregistration of the
		// Map.
		metricsMap, err = bpf.OpenMap(file)

		if err != nil {
			return fmt.Errorf("Unable to open metrics map: %s", err)
		}
	}

	var key, nextKey Key
	for {
		err := bpf.GetNextKey(metricsMap.GetFd(), unsafe.Pointer(&key), unsafe.Pointer(&nextKey))
		if err != nil {
			break
		}
		err = bpf.LookupElement(metricsMap.GetFd(), unsafe.Pointer(&nextKey), unsafe.Pointer(&entry[0]))
		if err != nil {
			return fmt.Errorf("unable to lookup metrics map: %s", err)
		}

		// cannot use `range entry` since, if the first value for a particular
		// CPU is zero, it never iterates over the next non-zero value.
		for i := 0; i < possibleCpus; i++ {
			// Increment Prometheus metrics here.
			updatePrometheusMetrics(&nextKey, &entry[i])
		}
		key = nextKey

	}
	return nil
}

func init() {
	possibleCpus = common.GetNumPossibleCPUs(log)

	vs := make(Values, possibleCpus)

	// Metrics is a mapping of all packet drops and forwards associated with
	// the node on ingress/egress direction
	Metrics = bpf.NewPerCPUHashMap(
		MapName,
		&Key{},
		int(unsafe.Sizeof(Key{})),
		&vs,
		int(unsafe.Sizeof(Value{})),
		possibleCpus,
		MaxEntries,
		0, 0,
		bpf.ConvertKeyValue,
	)
}
