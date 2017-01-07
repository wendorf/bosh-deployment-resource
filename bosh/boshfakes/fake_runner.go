// This file was generated by counterfeiter
package boshfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-deployment-resource/bosh"
)

type FakeRunner struct {
	ExecuteStub        func(commandOpts interface{}) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		commandOpts interface{}
	}
	executeReturns struct {
		result1 error
	}
	GetResultStub        func(commandOpts interface{}) ([]byte, error)
	getResultMutex       sync.RWMutex
	getResultArgsForCall []struct {
		commandOpts interface{}
	}
	getResultReturns struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRunner) Execute(commandOpts interface{}) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		commandOpts interface{}
	}{commandOpts})
	fake.recordInvocation("Execute", []interface{}{commandOpts})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(commandOpts)
	}
	return fake.executeReturns.result1
}

func (fake *FakeRunner) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeRunner) ExecuteArgsForCall(i int) interface{} {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].commandOpts
}

func (fake *FakeRunner) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunner) GetResult(commandOpts interface{}) ([]byte, error) {
	fake.getResultMutex.Lock()
	fake.getResultArgsForCall = append(fake.getResultArgsForCall, struct {
		commandOpts interface{}
	}{commandOpts})
	fake.recordInvocation("GetResult", []interface{}{commandOpts})
	fake.getResultMutex.Unlock()
	if fake.GetResultStub != nil {
		return fake.GetResultStub(commandOpts)
	}
	return fake.getResultReturns.result1, fake.getResultReturns.result2
}

func (fake *FakeRunner) GetResultCallCount() int {
	fake.getResultMutex.RLock()
	defer fake.getResultMutex.RUnlock()
	return len(fake.getResultArgsForCall)
}

func (fake *FakeRunner) GetResultArgsForCall(i int) interface{} {
	fake.getResultMutex.RLock()
	defer fake.getResultMutex.RUnlock()
	return fake.getResultArgsForCall[i].commandOpts
}

func (fake *FakeRunner) GetResultReturns(result1 []byte, result2 error) {
	fake.GetResultStub = nil
	fake.getResultReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.getResultMutex.RLock()
	defer fake.getResultMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRunner) recordInvocation(key string, args []interface{}) {
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

var _ bosh.Runner = new(FakeRunner)