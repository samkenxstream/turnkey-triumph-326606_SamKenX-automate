// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/compliance/reporting/stats/stats.proto

package stats

import policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.stats.v1.StatsService/ReadSummary", "compliance:reporting:stats:summary", "compliance:reportSummary:get", "POST", "/compliance/reporting/stats/summary", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.stats.v1.StatsService/ReadTrend", "compliance:reporting:stats:trend", "compliance:reportTrend:get", "POST", "/compliance/reporting/stats/trend", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.stats.v1.StatsService/ReadProfiles", "compliance:reporting:stats:profiles", "compliance:reportProfiles:get", "POST", "/compliance/reporting/stats/profiles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.stats.v1.StatsService/ReadFailures", "compliance:reporting:stats:failures", "compliance:reportFailures:get", "POST", "/compliance/reporting/stats/failures", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
}