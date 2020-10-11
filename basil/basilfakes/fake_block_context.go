// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"io"
	"sync"
	"time"

	"github.com/opsidian/basil/basil"
)

type FakeBlockContext struct {
	DeadlineStub        func() (time.Time, bool)
	deadlineMutex       sync.RWMutex
	deadlineArgsForCall []struct {
	}
	deadlineReturns struct {
		result1 time.Time
		result2 bool
	}
	deadlineReturnsOnCall map[int]struct {
		result1 time.Time
		result2 bool
	}
	DoneStub        func() <-chan struct{}
	doneMutex       sync.RWMutex
	doneArgsForCall []struct {
	}
	doneReturns struct {
		result1 <-chan struct{}
	}
	doneReturnsOnCall map[int]struct {
		result1 <-chan struct{}
	}
	ErrStub        func() error
	errMutex       sync.RWMutex
	errArgsForCall []struct {
	}
	errReturns struct {
		result1 error
	}
	errReturnsOnCall map[int]struct {
		result1 error
	}
	LoggerStub        func() basil.Logger
	loggerMutex       sync.RWMutex
	loggerArgsForCall []struct {
	}
	loggerReturns struct {
		result1 basil.Logger
	}
	loggerReturnsOnCall map[int]struct {
		result1 basil.Logger
	}
	PublishBlockStub        func(basil.Block, func() error) (bool, error)
	publishBlockMutex       sync.RWMutex
	publishBlockArgsForCall []struct {
		arg1 basil.Block
		arg2 func() error
	}
	publishBlockReturns struct {
		result1 bool
		result2 error
	}
	publishBlockReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	UserContextStub        func() interface{}
	userContextMutex       sync.RWMutex
	userContextArgsForCall []struct {
	}
	userContextReturns struct {
		result1 interface{}
	}
	userContextReturnsOnCall map[int]struct {
		result1 interface{}
	}
	ValueStub        func(interface{}) interface{}
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
		arg1 interface{}
	}
	valueReturns struct {
		result1 interface{}
	}
	valueReturnsOnCall map[int]struct {
		result1 interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockContext) Deadline() (time.Time, bool) {
	fake.deadlineMutex.Lock()
	ret, specificReturn := fake.deadlineReturnsOnCall[len(fake.deadlineArgsForCall)]
	fake.deadlineArgsForCall = append(fake.deadlineArgsForCall, struct {
	}{})
	fake.recordInvocation("Deadline", []interface{}{})
	fake.deadlineMutex.Unlock()
	if fake.DeadlineStub != nil {
		return fake.DeadlineStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deadlineReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockContext) DeadlineCallCount() int {
	fake.deadlineMutex.RLock()
	defer fake.deadlineMutex.RUnlock()
	return len(fake.deadlineArgsForCall)
}

func (fake *FakeBlockContext) DeadlineCalls(stub func() (time.Time, bool)) {
	fake.deadlineMutex.Lock()
	defer fake.deadlineMutex.Unlock()
	fake.DeadlineStub = stub
}

func (fake *FakeBlockContext) DeadlineReturns(result1 time.Time, result2 bool) {
	fake.deadlineMutex.Lock()
	defer fake.deadlineMutex.Unlock()
	fake.DeadlineStub = nil
	fake.deadlineReturns = struct {
		result1 time.Time
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockContext) DeadlineReturnsOnCall(i int, result1 time.Time, result2 bool) {
	fake.deadlineMutex.Lock()
	defer fake.deadlineMutex.Unlock()
	fake.DeadlineStub = nil
	if fake.deadlineReturnsOnCall == nil {
		fake.deadlineReturnsOnCall = make(map[int]struct {
			result1 time.Time
			result2 bool
		})
	}
	fake.deadlineReturnsOnCall[i] = struct {
		result1 time.Time
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockContext) Done() <-chan struct{} {
	fake.doneMutex.Lock()
	ret, specificReturn := fake.doneReturnsOnCall[len(fake.doneArgsForCall)]
	fake.doneArgsForCall = append(fake.doneArgsForCall, struct {
	}{})
	fake.recordInvocation("Done", []interface{}{})
	fake.doneMutex.Unlock()
	if fake.DoneStub != nil {
		return fake.DoneStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.doneReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContext) DoneCallCount() int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return len(fake.doneArgsForCall)
}

func (fake *FakeBlockContext) DoneCalls(stub func() <-chan struct{}) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = stub
}

func (fake *FakeBlockContext) DoneReturns(result1 <-chan struct{}) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	fake.doneReturns = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *FakeBlockContext) DoneReturnsOnCall(i int, result1 <-chan struct{}) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	if fake.doneReturnsOnCall == nil {
		fake.doneReturnsOnCall = make(map[int]struct {
			result1 <-chan struct{}
		})
	}
	fake.doneReturnsOnCall[i] = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *FakeBlockContext) Err() error {
	fake.errMutex.Lock()
	ret, specificReturn := fake.errReturnsOnCall[len(fake.errArgsForCall)]
	fake.errArgsForCall = append(fake.errArgsForCall, struct {
	}{})
	fake.recordInvocation("Err", []interface{}{})
	fake.errMutex.Unlock()
	if fake.ErrStub != nil {
		return fake.ErrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.errReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContext) ErrCallCount() int {
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	return len(fake.errArgsForCall)
}

func (fake *FakeBlockContext) ErrCalls(stub func() error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = stub
}

func (fake *FakeBlockContext) ErrReturns(result1 error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = nil
	fake.errReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockContext) ErrReturnsOnCall(i int, result1 error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = nil
	if fake.errReturnsOnCall == nil {
		fake.errReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.errReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockContext) Logger() basil.Logger {
	fake.loggerMutex.Lock()
	ret, specificReturn := fake.loggerReturnsOnCall[len(fake.loggerArgsForCall)]
	fake.loggerArgsForCall = append(fake.loggerArgsForCall, struct {
	}{})
	fake.recordInvocation("Logger", []interface{}{})
	fake.loggerMutex.Unlock()
	if fake.LoggerStub != nil {
		return fake.LoggerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.loggerReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContext) LoggerCallCount() int {
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	return len(fake.loggerArgsForCall)
}

func (fake *FakeBlockContext) LoggerCalls(stub func() basil.Logger) {
	fake.loggerMutex.Lock()
	defer fake.loggerMutex.Unlock()
	fake.LoggerStub = stub
}

func (fake *FakeBlockContext) LoggerReturns(result1 basil.Logger) {
	fake.loggerMutex.Lock()
	defer fake.loggerMutex.Unlock()
	fake.LoggerStub = nil
	fake.loggerReturns = struct {
		result1 basil.Logger
	}{result1}
}

func (fake *FakeBlockContext) LoggerReturnsOnCall(i int, result1 basil.Logger) {
	fake.loggerMutex.Lock()
	defer fake.loggerMutex.Unlock()
	fake.LoggerStub = nil
	if fake.loggerReturnsOnCall == nil {
		fake.loggerReturnsOnCall = make(map[int]struct {
			result1 basil.Logger
		})
	}
	fake.loggerReturnsOnCall[i] = struct {
		result1 basil.Logger
	}{result1}
}

func (fake *FakeBlockContext) PublishBlock(arg1 basil.Block, arg2 func() error) (bool, error) {
	fake.publishBlockMutex.Lock()
	ret, specificReturn := fake.publishBlockReturnsOnCall[len(fake.publishBlockArgsForCall)]
	fake.publishBlockArgsForCall = append(fake.publishBlockArgsForCall, struct {
		arg1 basil.Block
		arg2 func() error
	}{arg1, arg2})
	fake.recordInvocation("PublishBlock", []interface{}{arg1, arg2})
	fake.publishBlockMutex.Unlock()
	if fake.PublishBlockStub != nil {
		return fake.PublishBlockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.publishBlockReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockContext) PublishBlockCallCount() int {
	fake.publishBlockMutex.RLock()
	defer fake.publishBlockMutex.RUnlock()
	return len(fake.publishBlockArgsForCall)
}

func (fake *FakeBlockContext) PublishBlockCalls(stub func(basil.Block, func() error) (bool, error)) {
	fake.publishBlockMutex.Lock()
	defer fake.publishBlockMutex.Unlock()
	fake.PublishBlockStub = stub
}

func (fake *FakeBlockContext) PublishBlockArgsForCall(i int) (basil.Block, func() error) {
	fake.publishBlockMutex.RLock()
	defer fake.publishBlockMutex.RUnlock()
	argsForCall := fake.publishBlockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBlockContext) PublishBlockReturns(result1 bool, result2 error) {
	fake.publishBlockMutex.Lock()
	defer fake.publishBlockMutex.Unlock()
	fake.PublishBlockStub = nil
	fake.publishBlockReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBlockContext) PublishBlockReturnsOnCall(i int, result1 bool, result2 error) {
	fake.publishBlockMutex.Lock()
	defer fake.publishBlockMutex.Unlock()
	fake.PublishBlockStub = nil
	if fake.publishBlockReturnsOnCall == nil {
		fake.publishBlockReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.publishBlockReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBlockContext) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if fake.StdoutStub != nil {
		return fake.StdoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stdoutReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContext) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeBlockContext) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakeBlockContext) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeBlockContext) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeBlockContext) UserContext() interface{} {
	fake.userContextMutex.Lock()
	ret, specificReturn := fake.userContextReturnsOnCall[len(fake.userContextArgsForCall)]
	fake.userContextArgsForCall = append(fake.userContextArgsForCall, struct {
	}{})
	fake.recordInvocation("UserContext", []interface{}{})
	fake.userContextMutex.Unlock()
	if fake.UserContextStub != nil {
		return fake.UserContextStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.userContextReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContext) UserContextCallCount() int {
	fake.userContextMutex.RLock()
	defer fake.userContextMutex.RUnlock()
	return len(fake.userContextArgsForCall)
}

func (fake *FakeBlockContext) UserContextCalls(stub func() interface{}) {
	fake.userContextMutex.Lock()
	defer fake.userContextMutex.Unlock()
	fake.UserContextStub = stub
}

func (fake *FakeBlockContext) UserContextReturns(result1 interface{}) {
	fake.userContextMutex.Lock()
	defer fake.userContextMutex.Unlock()
	fake.UserContextStub = nil
	fake.userContextReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockContext) UserContextReturnsOnCall(i int, result1 interface{}) {
	fake.userContextMutex.Lock()
	defer fake.userContextMutex.Unlock()
	fake.UserContextStub = nil
	if fake.userContextReturnsOnCall == nil {
		fake.userContextReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.userContextReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockContext) Value(arg1 interface{}) interface{} {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Value", []interface{}{arg1})
	fake.valueMutex.Unlock()
	if fake.ValueStub != nil {
		return fake.ValueStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.valueReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContext) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeBlockContext) ValueCalls(stub func(interface{}) interface{}) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeBlockContext) ValueArgsForCall(i int) interface{} {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	argsForCall := fake.valueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockContext) ValueReturns(result1 interface{}) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockContext) ValueReturnsOnCall(i int, result1 interface{}) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deadlineMutex.RLock()
	defer fake.deadlineMutex.RUnlock()
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	fake.publishBlockMutex.RLock()
	defer fake.publishBlockMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.userContextMutex.RLock()
	defer fake.userContextMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockContext) recordInvocation(key string, args []interface{}) {
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

var _ basil.BlockContext = new(FakeBlockContext)
