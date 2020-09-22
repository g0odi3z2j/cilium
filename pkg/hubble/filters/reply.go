// Copyright 2019-2020 Authors of Hubble
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

package filters

import (
	"context"
	monitorAPI "github.com/cilium/cilium/pkg/monitor/api"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	v1 "github.com/cilium/cilium/pkg/hubble/api/v1"
)

func filterByReplyField(replyParams []bool) FilterFunc {
	return func(ev *v1.Event) bool {
		if len(replyParams) == 0 {
			return true
		}
		switch f := ev.Event.(type) {
		case v1.Flow:
			// FIMXE: In events originating from TraceNotify, the `reply`
			// field is determined from the connection tracking state in
			// the datapath. Unfortunately, not all trace points have the
			// connection tracking state available. For certain trace point
			// events, we do not know if it actually was a reply or not.
			//
			// Ideally, we would fix this in the parser and make the `reply`
			// an optional boolean, so we can distinguish between a `false`
			// value and an absent value. This however is a breaking change
			// in the API. Therefore, we only report flows here for which we
			// know that the reply field is reliable:
			if f.GetEventType().GetType() == monitorAPI.MessageTypeTrace {
				obsPoint := uint8(f.GetEventType().GetSubType())
				if !monitorAPI.TraceObservationPointHasConnState(obsPoint) {
					return false
				}
			}
			// For PolicyVerdict and Drop events, we statically assume they
			// always have `reply=false`, as the flow is not yet established
			// when the policy verdict is taken.
			// For all other events (including Accesslog/L7), we assume the
			// parser populates the Reply field reliably.

			reply := f.GetReply()
			for _, replyParam := range replyParams {
				if reply == replyParam {
					return true
				}
			}
		}
		return false
	}
}

// ReplyFilter implements filtering for reply flows
type ReplyFilter struct{}

// OnBuildFilter builds a reply filter
func (r *ReplyFilter) OnBuildFilter(ctx context.Context, ff *flowpb.FlowFilter) ([]FilterFunc, error) {
	var fs []FilterFunc

	if ff.GetReply() != nil {
		fs = append(fs, filterByReplyField(ff.GetReply()))
	}

	return fs, nil
}
