// Code generated by counterfeiter. DO NOT EDIT.
package testfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/test"
)

type FakeBlockWithInit struct {
	InitStub        func(basil.BlockContext) (bool, error)
	initMutex       sync.RWMutex
	initArgsForCall []struct {
		arg1 basil.BlockContext
	}
	initReturns struct {
		result1 bool
		result2 error
	}
	initReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockWithInit) Init(arg1 basil.BlockContext) (bool, error) {
	fake.initMutex.Lock()
	ret, specificReturn := fake.initReturnsOnCall[len(fake.initArgsForCall)]
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
		arg1 basil.BlockContext
	}{arg1})
	fake.recordInvocation("Init", []interface{}{arg1})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.initReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockWithInit) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeBlockWithInit) InitCalls(stub func(basil.BlockContext) (bool, error)) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = stub
}

func (fake *FakeBlockWithInit) InitArgsForCall(i int) basil.BlockContext {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	argsForCall := fake.initArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockWithInit) InitReturns(result1 bool, result2 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBlockWithInit) InitReturnsOnCall(i int, result1 bool, result2 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	if fake.initReturnsOnCall == nil {
		fake.initReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.initReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBlockWithInit) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockWithInit) recordInvocation(key string, args []interface{}) {
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

var _ test.BlockWithInit = new(FakeBlockWithInit)
