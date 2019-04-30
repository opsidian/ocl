// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
)

type FakeBlockNodeRegistry struct {
	AddBlockNodeStub        func(basil.BlockNode) error
	addBlockNodeMutex       sync.RWMutex
	addBlockNodeArgsForCall []struct {
		arg1 basil.BlockNode
	}
	addBlockNodeReturns struct {
		result1 error
	}
	addBlockNodeReturnsOnCall map[int]struct {
		result1 error
	}
	BlockNodeStub        func(basil.ID) (basil.BlockNode, bool)
	blockNodeMutex       sync.RWMutex
	blockNodeArgsForCall []struct {
		arg1 basil.ID
	}
	blockNodeReturns struct {
		result1 basil.BlockNode
		result2 bool
	}
	blockNodeReturnsOnCall map[int]struct {
		result1 basil.BlockNode
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockNodeRegistry) AddBlockNode(arg1 basil.BlockNode) error {
	fake.addBlockNodeMutex.Lock()
	ret, specificReturn := fake.addBlockNodeReturnsOnCall[len(fake.addBlockNodeArgsForCall)]
	fake.addBlockNodeArgsForCall = append(fake.addBlockNodeArgsForCall, struct {
		arg1 basil.BlockNode
	}{arg1})
	fake.recordInvocation("AddBlockNode", []interface{}{arg1})
	fake.addBlockNodeMutex.Unlock()
	if fake.AddBlockNodeStub != nil {
		return fake.AddBlockNodeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addBlockNodeReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNodeRegistry) AddBlockNodeCallCount() int {
	fake.addBlockNodeMutex.RLock()
	defer fake.addBlockNodeMutex.RUnlock()
	return len(fake.addBlockNodeArgsForCall)
}

func (fake *FakeBlockNodeRegistry) AddBlockNodeCalls(stub func(basil.BlockNode) error) {
	fake.addBlockNodeMutex.Lock()
	defer fake.addBlockNodeMutex.Unlock()
	fake.AddBlockNodeStub = stub
}

func (fake *FakeBlockNodeRegistry) AddBlockNodeArgsForCall(i int) basil.BlockNode {
	fake.addBlockNodeMutex.RLock()
	defer fake.addBlockNodeMutex.RUnlock()
	argsForCall := fake.addBlockNodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockNodeRegistry) AddBlockNodeReturns(result1 error) {
	fake.addBlockNodeMutex.Lock()
	defer fake.addBlockNodeMutex.Unlock()
	fake.AddBlockNodeStub = nil
	fake.addBlockNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockNodeRegistry) AddBlockNodeReturnsOnCall(i int, result1 error) {
	fake.addBlockNodeMutex.Lock()
	defer fake.addBlockNodeMutex.Unlock()
	fake.AddBlockNodeStub = nil
	if fake.addBlockNodeReturnsOnCall == nil {
		fake.addBlockNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addBlockNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockNodeRegistry) BlockNode(arg1 basil.ID) (basil.BlockNode, bool) {
	fake.blockNodeMutex.Lock()
	ret, specificReturn := fake.blockNodeReturnsOnCall[len(fake.blockNodeArgsForCall)]
	fake.blockNodeArgsForCall = append(fake.blockNodeArgsForCall, struct {
		arg1 basil.ID
	}{arg1})
	fake.recordInvocation("BlockNode", []interface{}{arg1})
	fake.blockNodeMutex.Unlock()
	if fake.BlockNodeStub != nil {
		return fake.BlockNodeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.blockNodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockNodeRegistry) BlockNodeCallCount() int {
	fake.blockNodeMutex.RLock()
	defer fake.blockNodeMutex.RUnlock()
	return len(fake.blockNodeArgsForCall)
}

func (fake *FakeBlockNodeRegistry) BlockNodeCalls(stub func(basil.ID) (basil.BlockNode, bool)) {
	fake.blockNodeMutex.Lock()
	defer fake.blockNodeMutex.Unlock()
	fake.BlockNodeStub = stub
}

func (fake *FakeBlockNodeRegistry) BlockNodeArgsForCall(i int) basil.ID {
	fake.blockNodeMutex.RLock()
	defer fake.blockNodeMutex.RUnlock()
	argsForCall := fake.blockNodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockNodeRegistry) BlockNodeReturns(result1 basil.BlockNode, result2 bool) {
	fake.blockNodeMutex.Lock()
	defer fake.blockNodeMutex.Unlock()
	fake.BlockNodeStub = nil
	fake.blockNodeReturns = struct {
		result1 basil.BlockNode
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockNodeRegistry) BlockNodeReturnsOnCall(i int, result1 basil.BlockNode, result2 bool) {
	fake.blockNodeMutex.Lock()
	defer fake.blockNodeMutex.Unlock()
	fake.BlockNodeStub = nil
	if fake.blockNodeReturnsOnCall == nil {
		fake.blockNodeReturnsOnCall = make(map[int]struct {
			result1 basil.BlockNode
			result2 bool
		})
	}
	fake.blockNodeReturnsOnCall[i] = struct {
		result1 basil.BlockNode
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockNodeRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addBlockNodeMutex.RLock()
	defer fake.addBlockNodeMutex.RUnlock()
	fake.blockNodeMutex.RLock()
	defer fake.blockNodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockNodeRegistry) recordInvocation(key string, args []interface{}) {
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

var _ basil.BlockNodeRegistry = new(FakeBlockNodeRegistry)