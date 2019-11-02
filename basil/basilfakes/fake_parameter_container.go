// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type FakeParameterContainer struct {
	BlockContainerStub        func() basil.BlockContainer
	blockContainerMutex       sync.RWMutex
	blockContainerArgsForCall []struct {
	}
	blockContainerReturns struct {
		result1 basil.BlockContainer
	}
	blockContainerReturnsOnCall map[int]struct {
		result1 basil.BlockContainer
	}
	CancelStub        func() bool
	cancelMutex       sync.RWMutex
	cancelArgsForCall []struct {
	}
	cancelReturns struct {
		result1 bool
	}
	cancelReturnsOnCall map[int]struct {
		result1 bool
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	JobIDStub        func() basil.ID
	jobIDMutex       sync.RWMutex
	jobIDArgsForCall []struct {
	}
	jobIDReturns struct {
		result1 basil.ID
	}
	jobIDReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	JobNameStub        func() basil.ID
	jobNameMutex       sync.RWMutex
	jobNameArgsForCall []struct {
	}
	jobNameReturns struct {
		result1 basil.ID
	}
	jobNameReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	LightweightStub        func() bool
	lightweightMutex       sync.RWMutex
	lightweightArgsForCall []struct {
	}
	lightweightReturns struct {
		result1 bool
	}
	lightweightReturnsOnCall map[int]struct {
		result1 bool
	}
	NodeStub        func() basil.Node
	nodeMutex       sync.RWMutex
	nodeArgsForCall []struct {
	}
	nodeReturns struct {
		result1 basil.Node
	}
	nodeReturnsOnCall map[int]struct {
		result1 basil.Node
	}
	RunStub        func()
	runMutex       sync.RWMutex
	runArgsForCall []struct {
	}
	SetJobIDStub        func(basil.ID)
	setJobIDMutex       sync.RWMutex
	setJobIDArgsForCall []struct {
		arg1 basil.ID
	}
	ValueStub        func() (interface{}, parsley.Error)
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
	}
	valueReturns struct {
		result1 interface{}
		result2 parsley.Error
	}
	valueReturnsOnCall map[int]struct {
		result1 interface{}
		result2 parsley.Error
	}
	WaitGroupsStub        func() []basil.WaitGroup
	waitGroupsMutex       sync.RWMutex
	waitGroupsArgsForCall []struct {
	}
	waitGroupsReturns struct {
		result1 []basil.WaitGroup
	}
	waitGroupsReturnsOnCall map[int]struct {
		result1 []basil.WaitGroup
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParameterContainer) BlockContainer() basil.BlockContainer {
	fake.blockContainerMutex.Lock()
	ret, specificReturn := fake.blockContainerReturnsOnCall[len(fake.blockContainerArgsForCall)]
	fake.blockContainerArgsForCall = append(fake.blockContainerArgsForCall, struct {
	}{})
	fake.recordInvocation("BlockContainer", []interface{}{})
	fake.blockContainerMutex.Unlock()
	if fake.BlockContainerStub != nil {
		return fake.BlockContainerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.blockContainerReturns
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) BlockContainerCallCount() int {
	fake.blockContainerMutex.RLock()
	defer fake.blockContainerMutex.RUnlock()
	return len(fake.blockContainerArgsForCall)
}

func (fake *FakeParameterContainer) BlockContainerCalls(stub func() basil.BlockContainer) {
	fake.blockContainerMutex.Lock()
	defer fake.blockContainerMutex.Unlock()
	fake.BlockContainerStub = stub
}

func (fake *FakeParameterContainer) BlockContainerReturns(result1 basil.BlockContainer) {
	fake.blockContainerMutex.Lock()
	defer fake.blockContainerMutex.Unlock()
	fake.BlockContainerStub = nil
	fake.blockContainerReturns = struct {
		result1 basil.BlockContainer
	}{result1}
}

func (fake *FakeParameterContainer) BlockContainerReturnsOnCall(i int, result1 basil.BlockContainer) {
	fake.blockContainerMutex.Lock()
	defer fake.blockContainerMutex.Unlock()
	fake.BlockContainerStub = nil
	if fake.blockContainerReturnsOnCall == nil {
		fake.blockContainerReturnsOnCall = make(map[int]struct {
			result1 basil.BlockContainer
		})
	}
	fake.blockContainerReturnsOnCall[i] = struct {
		result1 basil.BlockContainer
	}{result1}
}

func (fake *FakeParameterContainer) Cancel() bool {
	fake.cancelMutex.Lock()
	ret, specificReturn := fake.cancelReturnsOnCall[len(fake.cancelArgsForCall)]
	fake.cancelArgsForCall = append(fake.cancelArgsForCall, struct {
	}{})
	fake.recordInvocation("Cancel", []interface{}{})
	fake.cancelMutex.Unlock()
	if fake.CancelStub != nil {
		return fake.CancelStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cancelReturns
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) CancelCallCount() int {
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	return len(fake.cancelArgsForCall)
}

func (fake *FakeParameterContainer) CancelCalls(stub func() bool) {
	fake.cancelMutex.Lock()
	defer fake.cancelMutex.Unlock()
	fake.CancelStub = stub
}

func (fake *FakeParameterContainer) CancelReturns(result1 bool) {
	fake.cancelMutex.Lock()
	defer fake.cancelMutex.Unlock()
	fake.CancelStub = nil
	fake.cancelReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterContainer) CancelReturnsOnCall(i int, result1 bool) {
	fake.cancelMutex.Lock()
	defer fake.cancelMutex.Unlock()
	fake.CancelStub = nil
	if fake.cancelReturnsOnCall == nil {
		fake.cancelReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.cancelReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterContainer) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeParameterContainer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeParameterContainer) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeParameterContainer) JobID() basil.ID {
	fake.jobIDMutex.Lock()
	ret, specificReturn := fake.jobIDReturnsOnCall[len(fake.jobIDArgsForCall)]
	fake.jobIDArgsForCall = append(fake.jobIDArgsForCall, struct {
	}{})
	fake.recordInvocation("JobID", []interface{}{})
	fake.jobIDMutex.Unlock()
	if fake.JobIDStub != nil {
		return fake.JobIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.jobIDReturns
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) JobIDCallCount() int {
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	return len(fake.jobIDArgsForCall)
}

func (fake *FakeParameterContainer) JobIDCalls(stub func() basil.ID) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = stub
}

func (fake *FakeParameterContainer) JobIDReturns(result1 basil.ID) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = nil
	fake.jobIDReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeParameterContainer) JobIDReturnsOnCall(i int, result1 basil.ID) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = nil
	if fake.jobIDReturnsOnCall == nil {
		fake.jobIDReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.jobIDReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeParameterContainer) JobName() basil.ID {
	fake.jobNameMutex.Lock()
	ret, specificReturn := fake.jobNameReturnsOnCall[len(fake.jobNameArgsForCall)]
	fake.jobNameArgsForCall = append(fake.jobNameArgsForCall, struct {
	}{})
	fake.recordInvocation("JobName", []interface{}{})
	fake.jobNameMutex.Unlock()
	if fake.JobNameStub != nil {
		return fake.JobNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.jobNameReturns
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) JobNameCallCount() int {
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	return len(fake.jobNameArgsForCall)
}

func (fake *FakeParameterContainer) JobNameCalls(stub func() basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = stub
}

func (fake *FakeParameterContainer) JobNameReturns(result1 basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = nil
	fake.jobNameReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeParameterContainer) JobNameReturnsOnCall(i int, result1 basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = nil
	if fake.jobNameReturnsOnCall == nil {
		fake.jobNameReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.jobNameReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeParameterContainer) Lightweight() bool {
	fake.lightweightMutex.Lock()
	ret, specificReturn := fake.lightweightReturnsOnCall[len(fake.lightweightArgsForCall)]
	fake.lightweightArgsForCall = append(fake.lightweightArgsForCall, struct {
	}{})
	fake.recordInvocation("Lightweight", []interface{}{})
	fake.lightweightMutex.Unlock()
	if fake.LightweightStub != nil {
		return fake.LightweightStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lightweightReturns
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) LightweightCallCount() int {
	fake.lightweightMutex.RLock()
	defer fake.lightweightMutex.RUnlock()
	return len(fake.lightweightArgsForCall)
}

func (fake *FakeParameterContainer) LightweightCalls(stub func() bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = stub
}

func (fake *FakeParameterContainer) LightweightReturns(result1 bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = nil
	fake.lightweightReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterContainer) LightweightReturnsOnCall(i int, result1 bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = nil
	if fake.lightweightReturnsOnCall == nil {
		fake.lightweightReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.lightweightReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterContainer) Node() basil.Node {
	fake.nodeMutex.Lock()
	ret, specificReturn := fake.nodeReturnsOnCall[len(fake.nodeArgsForCall)]
	fake.nodeArgsForCall = append(fake.nodeArgsForCall, struct {
	}{})
	fake.recordInvocation("Node", []interface{}{})
	fake.nodeMutex.Unlock()
	if fake.NodeStub != nil {
		return fake.NodeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nodeReturns
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) NodeCallCount() int {
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	return len(fake.nodeArgsForCall)
}

func (fake *FakeParameterContainer) NodeCalls(stub func() basil.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = stub
}

func (fake *FakeParameterContainer) NodeReturns(result1 basil.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = nil
	fake.nodeReturns = struct {
		result1 basil.Node
	}{result1}
}

func (fake *FakeParameterContainer) NodeReturnsOnCall(i int, result1 basil.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = nil
	if fake.nodeReturnsOnCall == nil {
		fake.nodeReturnsOnCall = make(map[int]struct {
			result1 basil.Node
		})
	}
	fake.nodeReturnsOnCall[i] = struct {
		result1 basil.Node
	}{result1}
}

func (fake *FakeParameterContainer) Run() {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
	}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub()
	}
}

func (fake *FakeParameterContainer) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeParameterContainer) RunCalls(stub func()) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeParameterContainer) SetJobID(arg1 basil.ID) {
	fake.setJobIDMutex.Lock()
	fake.setJobIDArgsForCall = append(fake.setJobIDArgsForCall, struct {
		arg1 basil.ID
	}{arg1})
	fake.recordInvocation("SetJobID", []interface{}{arg1})
	fake.setJobIDMutex.Unlock()
	if fake.SetJobIDStub != nil {
		fake.SetJobIDStub(arg1)
	}
}

func (fake *FakeParameterContainer) SetJobIDCallCount() int {
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	return len(fake.setJobIDArgsForCall)
}

func (fake *FakeParameterContainer) SetJobIDCalls(stub func(basil.ID)) {
	fake.setJobIDMutex.Lock()
	defer fake.setJobIDMutex.Unlock()
	fake.SetJobIDStub = stub
}

func (fake *FakeParameterContainer) SetJobIDArgsForCall(i int) basil.ID {
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	argsForCall := fake.setJobIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParameterContainer) Value() (interface{}, parsley.Error) {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
	}{})
	fake.recordInvocation("Value", []interface{}{})
	fake.valueMutex.Unlock()
	if fake.ValueStub != nil {
		return fake.ValueStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.valueReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParameterContainer) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeParameterContainer) ValueCalls(stub func() (interface{}, parsley.Error)) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeParameterContainer) ValueReturns(result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeParameterContainer) ValueReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 parsley.Error
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeParameterContainer) WaitGroups() []basil.WaitGroup {
	fake.waitGroupsMutex.Lock()
	ret, specificReturn := fake.waitGroupsReturnsOnCall[len(fake.waitGroupsArgsForCall)]
	fake.waitGroupsArgsForCall = append(fake.waitGroupsArgsForCall, struct {
	}{})
	fake.recordInvocation("WaitGroups", []interface{}{})
	fake.waitGroupsMutex.Unlock()
	if fake.WaitGroupsStub != nil {
		return fake.WaitGroupsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitGroupsReturns
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) WaitGroupsCallCount() int {
	fake.waitGroupsMutex.RLock()
	defer fake.waitGroupsMutex.RUnlock()
	return len(fake.waitGroupsArgsForCall)
}

func (fake *FakeParameterContainer) WaitGroupsCalls(stub func() []basil.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = stub
}

func (fake *FakeParameterContainer) WaitGroupsReturns(result1 []basil.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = nil
	fake.waitGroupsReturns = struct {
		result1 []basil.WaitGroup
	}{result1}
}

func (fake *FakeParameterContainer) WaitGroupsReturnsOnCall(i int, result1 []basil.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = nil
	if fake.waitGroupsReturnsOnCall == nil {
		fake.waitGroupsReturnsOnCall = make(map[int]struct {
			result1 []basil.WaitGroup
		})
	}
	fake.waitGroupsReturnsOnCall[i] = struct {
		result1 []basil.WaitGroup
	}{result1}
}

func (fake *FakeParameterContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockContainerMutex.RLock()
	defer fake.blockContainerMutex.RUnlock()
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	fake.lightweightMutex.RLock()
	defer fake.lightweightMutex.RUnlock()
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	fake.waitGroupsMutex.RLock()
	defer fake.waitGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParameterContainer) recordInvocation(key string, args []interface{}) {
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

var _ basil.ParameterContainer = new(FakeParameterContainer)
