// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ConsoleStatusApplyConfiguration represents an declarative configuration of the ConsoleStatus type for use
// with apply.
type ConsoleStatusApplyConfiguration struct {
	OperatorStatusApplyConfiguration `json:",inline"`
}

// ConsoleStatusApplyConfiguration constructs an declarative configuration of the ConsoleStatus type for use with
// apply.
func ConsoleStatus() *ConsoleStatusApplyConfiguration {
	return &ConsoleStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ConsoleStatusApplyConfiguration) WithObservedGeneration(value int64) *ConsoleStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ConsoleStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *ConsoleStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ConsoleStatusApplyConfiguration) WithVersion(value string) *ConsoleStatusApplyConfiguration {
	b.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *ConsoleStatusApplyConfiguration) WithReadyReplicas(value int32) *ConsoleStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *ConsoleStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *ConsoleStatusApplyConfiguration {
	b.LatestAvailableRevision = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *ConsoleStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *ConsoleStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.Generations = append(b.Generations, *values[i])
	}
	return b
}
