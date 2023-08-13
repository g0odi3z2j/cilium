// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package translation

import (
	"sort"
	"testing"

	envoy_config_route_v3 "github.com/cilium/proxy/go/envoy/config/route/v3"
	envoy_type_matcher_v3 "github.com/cilium/proxy/go/envoy/type/matcher/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cilium/cilium/operator/pkg/model"
)

func TestSortableRoute(t *testing.T) {
	arr := SortableRoute{
		{
			Name: "exact match 1",
			Match: &envoy_config_route_v3.RouteMatch{
				PathSpecifier: &envoy_config_route_v3.RouteMatch_Path{
					Path: "/exact/match",
				},
			},
		},
		{
			Name: "another exact match",
			Match: &envoy_config_route_v3.RouteMatch{
				PathSpecifier: &envoy_config_route_v3.RouteMatch_Path{
					Path: "/exact/match/another",
				},
			},
		},
		{
			Name: "prefix match",
			Match: &envoy_config_route_v3.RouteMatch{
				PathSpecifier: &envoy_config_route_v3.RouteMatch_SafeRegex{
					SafeRegex: &envoy_type_matcher_v3.RegexMatcher{
						Regex: "/prefix/match",
					},
				},
			},
		},
		{
			Name: "another prefix match",
			Match: &envoy_config_route_v3.RouteMatch{
				PathSpecifier: &envoy_config_route_v3.RouteMatch_SafeRegex{
					SafeRegex: &envoy_type_matcher_v3.RegexMatcher{
						Regex: "/prefix/match/another",
					},
				},
			},
		},
		{
			Name: "prefix match with one header match",
			Match: &envoy_config_route_v3.RouteMatch{
				PathSpecifier: &envoy_config_route_v3.RouteMatch_SafeRegex{
					SafeRegex: &envoy_type_matcher_v3.RegexMatcher{
						Regex: "/header",
					},
				},
				Headers: []*envoy_config_route_v3.HeaderMatcher{
					{
						Name: "header1",
						HeaderMatchSpecifier: &envoy_config_route_v3.HeaderMatcher_StringMatch{
							StringMatch: &envoy_type_matcher_v3.StringMatcher{
								MatchPattern: &envoy_type_matcher_v3.StringMatcher_Exact{
									Exact: "value1",
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "prefix match with two header matches",
			Match: &envoy_config_route_v3.RouteMatch{
				PathSpecifier: &envoy_config_route_v3.RouteMatch_SafeRegex{
					SafeRegex: &envoy_type_matcher_v3.RegexMatcher{
						Regex: "/header",
					},
				},
				Headers: []*envoy_config_route_v3.HeaderMatcher{
					{
						Name: "header1",
						HeaderMatchSpecifier: &envoy_config_route_v3.HeaderMatcher_StringMatch{
							StringMatch: &envoy_type_matcher_v3.StringMatcher{
								MatchPattern: &envoy_type_matcher_v3.StringMatcher_Exact{
									Exact: "value1",
								},
							},
						},
					},
					{
						Name: "header2",
						HeaderMatchSpecifier: &envoy_config_route_v3.HeaderMatcher_StringMatch{
							StringMatch: &envoy_type_matcher_v3.StringMatcher{
								MatchPattern: &envoy_type_matcher_v3.StringMatcher_Exact{
									Exact: "value2",
								},
							},
						},
					},
				},
			},
		},
	}

	sort.Sort(arr)

	// Exact match comes first in any order
	assert.True(t, len(arr[0].Match.GetPath()) != 0)
	assert.True(t, len(arr[1].Match.GetPath()) != 0)

	// Prefix match with longer path comes first
	assert.Equal(t, "/prefix/match/another", arr[2].Match.GetSafeRegex().GetRegex())
	assert.Equal(t, "/prefix/match", arr[3].Match.GetSafeRegex().GetRegex())

	// More Header match comes first
	assert.True(t, len(arr[4].Match.GetHeaders()) == 2)
	assert.True(t, len(arr[5].Match.GetHeaders()) == 1)
}

func Test_hostRewriteMutation(t *testing.T) {
	t.Run("no host rewrite", func(t *testing.T) {
		route := &envoy_config_route_v3.Route_Route{
			Route: &envoy_config_route_v3.RouteAction{},
		}
		res := hostRewriteMutation(nil)(route)
		require.Equal(t, route, res)
	})

	t.Run("with host rewrite", func(t *testing.T) {
		route := &envoy_config_route_v3.Route_Route{
			Route: &envoy_config_route_v3.RouteAction{},
		}
		rewrite := &model.HTTPURLRewriteFilter{
			HostName: model.AddressOf("example.com"),
		}

		res := hostRewriteMutation(rewrite)(route)
		require.Equal(t, res.Route.HostRewriteSpecifier, &envoy_config_route_v3.RouteAction_HostRewriteLiteral{
			HostRewriteLiteral: "example.com",
		})
	})
}

func Test_pathPrefixMutation(t *testing.T) {
	t.Run("no prefix rewrite", func(t *testing.T) {
		route := &envoy_config_route_v3.Route_Route{
			Route: &envoy_config_route_v3.RouteAction{},
		}
		res := pathPrefixMutation(nil)(route)
		require.Equal(t, route, res)
	})

	t.Run("with prefix rewrite", func(t *testing.T) {
		route := &envoy_config_route_v3.Route_Route{
			Route: &envoy_config_route_v3.RouteAction{},
		}
		rewrite := &model.HTTPURLRewriteFilter{
			Path: &model.StringMatch{
				Prefix: "/prefix",
			},
		}

		res := pathPrefixMutation(rewrite)(route)
		require.Equal(t, res.Route.PrefixRewrite, "/prefix")
	})
}
