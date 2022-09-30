// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PromQLClusterConditionApplyConfiguration represents an declarative configuration of the PromQLClusterCondition type for use
// with apply.
type PromQLClusterConditionApplyConfiguration struct {
	PromQL *string `json:"promql,omitempty"`
}

// PromQLClusterConditionApplyConfiguration constructs an declarative configuration of the PromQLClusterCondition type for use with
// apply.
func PromQLClusterCondition() *PromQLClusterConditionApplyConfiguration {
	return &PromQLClusterConditionApplyConfiguration{}
}

// WithPromQL sets the PromQL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PromQL field is set to the value of the last call.
func (b *PromQLClusterConditionApplyConfiguration) WithPromQL(value string) *PromQLClusterConditionApplyConfiguration {
	b.PromQL = &value
	return b
}