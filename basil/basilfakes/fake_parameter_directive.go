// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
)

type FakeParameterDirective struct {
	ApplyToParameterConfigStub        func(*basil.ParameterConfig)
	applyToParameterConfigMutex       sync.RWMutex
	applyToParameterConfigArgsForCall []struct {
		arg1 *basil.ParameterConfig
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParameterDirective) ApplyToParameterConfig(arg1 *basil.ParameterConfig) {
	fake.applyToParameterConfigMutex.Lock()
	fake.applyToParameterConfigArgsForCall = append(fake.applyToParameterConfigArgsForCall, struct {
		arg1 *basil.ParameterConfig
	}{arg1})
	fake.recordInvocation("ApplyToParameterConfig", []interface{}{arg1})
	fake.applyToParameterConfigMutex.Unlock()
	if fake.ApplyToParameterConfigStub != nil {
		fake.ApplyToParameterConfigStub(arg1)
	}
}

func (fake *FakeParameterDirective) ApplyToParameterConfigCallCount() int {
	fake.applyToParameterConfigMutex.RLock()
	defer fake.applyToParameterConfigMutex.RUnlock()
	return len(fake.applyToParameterConfigArgsForCall)
}

func (fake *FakeParameterDirective) ApplyToParameterConfigCalls(stub func(*basil.ParameterConfig)) {
	fake.applyToParameterConfigMutex.Lock()
	defer fake.applyToParameterConfigMutex.Unlock()
	fake.ApplyToParameterConfigStub = stub
}

func (fake *FakeParameterDirective) ApplyToParameterConfigArgsForCall(i int) *basil.ParameterConfig {
	fake.applyToParameterConfigMutex.RLock()
	defer fake.applyToParameterConfigMutex.RUnlock()
	argsForCall := fake.applyToParameterConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParameterDirective) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyToParameterConfigMutex.RLock()
	defer fake.applyToParameterConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParameterDirective) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ basil.ParameterDirective = new(FakeParameterDirective)
