// Code generated by counterfeiter. DO NOT EDIT.
package agentfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/agent"
)

type FakeStartManager struct {
	CanStartStub        func() bool
	canStartMutex       sync.RWMutex
	canStartArgsForCall []struct {
	}
	canStartReturns struct {
		result1 bool
	}
	canStartReturnsOnCall map[int]struct {
		result1 bool
	}
	RegisterStartStub        func() error
	registerStartMutex       sync.RWMutex
	registerStartArgsForCall []struct {
	}
	registerStartReturns struct {
		result1 error
	}
	registerStartReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStartManager) CanStart() bool {
	fake.canStartMutex.Lock()
	ret, specificReturn := fake.canStartReturnsOnCall[len(fake.canStartArgsForCall)]
	fake.canStartArgsForCall = append(fake.canStartArgsForCall, struct {
	}{})
	fake.recordInvocation("CanStart", []interface{}{})
	fake.canStartMutex.Unlock()
	if fake.CanStartStub != nil {
		return fake.CanStartStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.canStartReturns
	return fakeReturns.result1
}

func (fake *FakeStartManager) CanStartCallCount() int {
	fake.canStartMutex.RLock()
	defer fake.canStartMutex.RUnlock()
	return len(fake.canStartArgsForCall)
}

func (fake *FakeStartManager) CanStartCalls(stub func() bool) {
	fake.canStartMutex.Lock()
	defer fake.canStartMutex.Unlock()
	fake.CanStartStub = stub
}

func (fake *FakeStartManager) CanStartReturns(result1 bool) {
	fake.canStartMutex.Lock()
	defer fake.canStartMutex.Unlock()
	fake.CanStartStub = nil
	fake.canStartReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStartManager) CanStartReturnsOnCall(i int, result1 bool) {
	fake.canStartMutex.Lock()
	defer fake.canStartMutex.Unlock()
	fake.CanStartStub = nil
	if fake.canStartReturnsOnCall == nil {
		fake.canStartReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.canStartReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStartManager) RegisterStart() error {
	fake.registerStartMutex.Lock()
	ret, specificReturn := fake.registerStartReturnsOnCall[len(fake.registerStartArgsForCall)]
	fake.registerStartArgsForCall = append(fake.registerStartArgsForCall, struct {
	}{})
	fake.recordInvocation("RegisterStart", []interface{}{})
	fake.registerStartMutex.Unlock()
	if fake.RegisterStartStub != nil {
		return fake.RegisterStartStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.registerStartReturns
	return fakeReturns.result1
}

func (fake *FakeStartManager) RegisterStartCallCount() int {
	fake.registerStartMutex.RLock()
	defer fake.registerStartMutex.RUnlock()
	return len(fake.registerStartArgsForCall)
}

func (fake *FakeStartManager) RegisterStartCalls(stub func() error) {
	fake.registerStartMutex.Lock()
	defer fake.registerStartMutex.Unlock()
	fake.RegisterStartStub = stub
}

func (fake *FakeStartManager) RegisterStartReturns(result1 error) {
	fake.registerStartMutex.Lock()
	defer fake.registerStartMutex.Unlock()
	fake.RegisterStartStub = nil
	fake.registerStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStartManager) RegisterStartReturnsOnCall(i int, result1 error) {
	fake.registerStartMutex.Lock()
	defer fake.registerStartMutex.Unlock()
	fake.RegisterStartStub = nil
	if fake.registerStartReturnsOnCall == nil {
		fake.registerStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStartManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.canStartMutex.RLock()
	defer fake.canStartMutex.RUnlock()
	fake.registerStartMutex.RLock()
	defer fake.registerStartMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStartManager) recordInvocation(key string, args []interface{}) {
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

var _ agent.StartManager = new(FakeStartManager)
