// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/pivotal-michael-stergianis/cf-mgmt/organizationreader"
)

type FakeReader struct {
	AddOrgToListStub        func(cfclient.Org)
	addOrgToListMutex       sync.RWMutex
	addOrgToListArgsForCall []struct {
		arg1 cfclient.Org
	}
	ClearOrgListStub        func()
	clearOrgListMutex       sync.RWMutex
	clearOrgListArgsForCall []struct {
	}
	FindOrgStub        func(string) (cfclient.Org, error)
	findOrgMutex       sync.RWMutex
	findOrgArgsForCall []struct {
		arg1 string
	}
	findOrgReturns struct {
		result1 cfclient.Org
		result2 error
	}
	findOrgReturnsOnCall map[int]struct {
		result1 cfclient.Org
		result2 error
	}
	FindOrgByGUIDStub        func(string) (cfclient.Org, error)
	findOrgByGUIDMutex       sync.RWMutex
	findOrgByGUIDArgsForCall []struct {
		arg1 string
	}
	findOrgByGUIDReturns struct {
		result1 cfclient.Org
		result2 error
	}
	findOrgByGUIDReturnsOnCall map[int]struct {
		result1 cfclient.Org
		result2 error
	}
	GetOrgByGUIDStub        func(string) (cfclient.Org, error)
	getOrgByGUIDMutex       sync.RWMutex
	getOrgByGUIDArgsForCall []struct {
		arg1 string
	}
	getOrgByGUIDReturns struct {
		result1 cfclient.Org
		result2 error
	}
	getOrgByGUIDReturnsOnCall map[int]struct {
		result1 cfclient.Org
		result2 error
	}
	GetOrgGUIDStub        func(string) (string, error)
	getOrgGUIDMutex       sync.RWMutex
	getOrgGUIDArgsForCall []struct {
		arg1 string
	}
	getOrgGUIDReturns struct {
		result1 string
		result2 error
	}
	getOrgGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListOrgsStub        func() ([]cfclient.Org, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
	}
	listOrgsReturns struct {
		result1 []cfclient.Org
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []cfclient.Org
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReader) AddOrgToList(arg1 cfclient.Org) {
	fake.addOrgToListMutex.Lock()
	fake.addOrgToListArgsForCall = append(fake.addOrgToListArgsForCall, struct {
		arg1 cfclient.Org
	}{arg1})
	fake.recordInvocation("AddOrgToList", []interface{}{arg1})
	fake.addOrgToListMutex.Unlock()
	if fake.AddOrgToListStub != nil {
		fake.AddOrgToListStub(arg1)
	}
}

func (fake *FakeReader) AddOrgToListCallCount() int {
	fake.addOrgToListMutex.RLock()
	defer fake.addOrgToListMutex.RUnlock()
	return len(fake.addOrgToListArgsForCall)
}

func (fake *FakeReader) AddOrgToListCalls(stub func(cfclient.Org)) {
	fake.addOrgToListMutex.Lock()
	defer fake.addOrgToListMutex.Unlock()
	fake.AddOrgToListStub = stub
}

func (fake *FakeReader) AddOrgToListArgsForCall(i int) cfclient.Org {
	fake.addOrgToListMutex.RLock()
	defer fake.addOrgToListMutex.RUnlock()
	argsForCall := fake.addOrgToListArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) ClearOrgList() {
	fake.clearOrgListMutex.Lock()
	fake.clearOrgListArgsForCall = append(fake.clearOrgListArgsForCall, struct {
	}{})
	fake.recordInvocation("ClearOrgList", []interface{}{})
	fake.clearOrgListMutex.Unlock()
	if fake.ClearOrgListStub != nil {
		fake.ClearOrgListStub()
	}
}

func (fake *FakeReader) ClearOrgListCallCount() int {
	fake.clearOrgListMutex.RLock()
	defer fake.clearOrgListMutex.RUnlock()
	return len(fake.clearOrgListArgsForCall)
}

func (fake *FakeReader) ClearOrgListCalls(stub func()) {
	fake.clearOrgListMutex.Lock()
	defer fake.clearOrgListMutex.Unlock()
	fake.ClearOrgListStub = stub
}

func (fake *FakeReader) FindOrg(arg1 string) (cfclient.Org, error) {
	fake.findOrgMutex.Lock()
	ret, specificReturn := fake.findOrgReturnsOnCall[len(fake.findOrgArgsForCall)]
	fake.findOrgArgsForCall = append(fake.findOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindOrg", []interface{}{arg1})
	fake.findOrgMutex.Unlock()
	if fake.FindOrgStub != nil {
		return fake.FindOrgStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrgReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) FindOrgCallCount() int {
	fake.findOrgMutex.RLock()
	defer fake.findOrgMutex.RUnlock()
	return len(fake.findOrgArgsForCall)
}

func (fake *FakeReader) FindOrgCalls(stub func(string) (cfclient.Org, error)) {
	fake.findOrgMutex.Lock()
	defer fake.findOrgMutex.Unlock()
	fake.FindOrgStub = stub
}

func (fake *FakeReader) FindOrgArgsForCall(i int) string {
	fake.findOrgMutex.RLock()
	defer fake.findOrgMutex.RUnlock()
	argsForCall := fake.findOrgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) FindOrgReturns(result1 cfclient.Org, result2 error) {
	fake.findOrgMutex.Lock()
	defer fake.findOrgMutex.Unlock()
	fake.FindOrgStub = nil
	fake.findOrgReturns = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) FindOrgReturnsOnCall(i int, result1 cfclient.Org, result2 error) {
	fake.findOrgMutex.Lock()
	defer fake.findOrgMutex.Unlock()
	fake.FindOrgStub = nil
	if fake.findOrgReturnsOnCall == nil {
		fake.findOrgReturnsOnCall = make(map[int]struct {
			result1 cfclient.Org
			result2 error
		})
	}
	fake.findOrgReturnsOnCall[i] = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) FindOrgByGUID(arg1 string) (cfclient.Org, error) {
	fake.findOrgByGUIDMutex.Lock()
	ret, specificReturn := fake.findOrgByGUIDReturnsOnCall[len(fake.findOrgByGUIDArgsForCall)]
	fake.findOrgByGUIDArgsForCall = append(fake.findOrgByGUIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindOrgByGUID", []interface{}{arg1})
	fake.findOrgByGUIDMutex.Unlock()
	if fake.FindOrgByGUIDStub != nil {
		return fake.FindOrgByGUIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrgByGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) FindOrgByGUIDCallCount() int {
	fake.findOrgByGUIDMutex.RLock()
	defer fake.findOrgByGUIDMutex.RUnlock()
	return len(fake.findOrgByGUIDArgsForCall)
}

func (fake *FakeReader) FindOrgByGUIDCalls(stub func(string) (cfclient.Org, error)) {
	fake.findOrgByGUIDMutex.Lock()
	defer fake.findOrgByGUIDMutex.Unlock()
	fake.FindOrgByGUIDStub = stub
}

func (fake *FakeReader) FindOrgByGUIDArgsForCall(i int) string {
	fake.findOrgByGUIDMutex.RLock()
	defer fake.findOrgByGUIDMutex.RUnlock()
	argsForCall := fake.findOrgByGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) FindOrgByGUIDReturns(result1 cfclient.Org, result2 error) {
	fake.findOrgByGUIDMutex.Lock()
	defer fake.findOrgByGUIDMutex.Unlock()
	fake.FindOrgByGUIDStub = nil
	fake.findOrgByGUIDReturns = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) FindOrgByGUIDReturnsOnCall(i int, result1 cfclient.Org, result2 error) {
	fake.findOrgByGUIDMutex.Lock()
	defer fake.findOrgByGUIDMutex.Unlock()
	fake.FindOrgByGUIDStub = nil
	if fake.findOrgByGUIDReturnsOnCall == nil {
		fake.findOrgByGUIDReturnsOnCall = make(map[int]struct {
			result1 cfclient.Org
			result2 error
		})
	}
	fake.findOrgByGUIDReturnsOnCall[i] = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgByGUID(arg1 string) (cfclient.Org, error) {
	fake.getOrgByGUIDMutex.Lock()
	ret, specificReturn := fake.getOrgByGUIDReturnsOnCall[len(fake.getOrgByGUIDArgsForCall)]
	fake.getOrgByGUIDArgsForCall = append(fake.getOrgByGUIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrgByGUID", []interface{}{arg1})
	fake.getOrgByGUIDMutex.Unlock()
	if fake.GetOrgByGUIDStub != nil {
		return fake.GetOrgByGUIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getOrgByGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) GetOrgByGUIDCallCount() int {
	fake.getOrgByGUIDMutex.RLock()
	defer fake.getOrgByGUIDMutex.RUnlock()
	return len(fake.getOrgByGUIDArgsForCall)
}

func (fake *FakeReader) GetOrgByGUIDCalls(stub func(string) (cfclient.Org, error)) {
	fake.getOrgByGUIDMutex.Lock()
	defer fake.getOrgByGUIDMutex.Unlock()
	fake.GetOrgByGUIDStub = stub
}

func (fake *FakeReader) GetOrgByGUIDArgsForCall(i int) string {
	fake.getOrgByGUIDMutex.RLock()
	defer fake.getOrgByGUIDMutex.RUnlock()
	argsForCall := fake.getOrgByGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) GetOrgByGUIDReturns(result1 cfclient.Org, result2 error) {
	fake.getOrgByGUIDMutex.Lock()
	defer fake.getOrgByGUIDMutex.Unlock()
	fake.GetOrgByGUIDStub = nil
	fake.getOrgByGUIDReturns = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgByGUIDReturnsOnCall(i int, result1 cfclient.Org, result2 error) {
	fake.getOrgByGUIDMutex.Lock()
	defer fake.getOrgByGUIDMutex.Unlock()
	fake.GetOrgByGUIDStub = nil
	if fake.getOrgByGUIDReturnsOnCall == nil {
		fake.getOrgByGUIDReturnsOnCall = make(map[int]struct {
			result1 cfclient.Org
			result2 error
		})
	}
	fake.getOrgByGUIDReturnsOnCall[i] = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgGUID(arg1 string) (string, error) {
	fake.getOrgGUIDMutex.Lock()
	ret, specificReturn := fake.getOrgGUIDReturnsOnCall[len(fake.getOrgGUIDArgsForCall)]
	fake.getOrgGUIDArgsForCall = append(fake.getOrgGUIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrgGUID", []interface{}{arg1})
	fake.getOrgGUIDMutex.Unlock()
	if fake.GetOrgGUIDStub != nil {
		return fake.GetOrgGUIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getOrgGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) GetOrgGUIDCallCount() int {
	fake.getOrgGUIDMutex.RLock()
	defer fake.getOrgGUIDMutex.RUnlock()
	return len(fake.getOrgGUIDArgsForCall)
}

func (fake *FakeReader) GetOrgGUIDCalls(stub func(string) (string, error)) {
	fake.getOrgGUIDMutex.Lock()
	defer fake.getOrgGUIDMutex.Unlock()
	fake.GetOrgGUIDStub = stub
}

func (fake *FakeReader) GetOrgGUIDArgsForCall(i int) string {
	fake.getOrgGUIDMutex.RLock()
	defer fake.getOrgGUIDMutex.RUnlock()
	argsForCall := fake.getOrgGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) GetOrgGUIDReturns(result1 string, result2 error) {
	fake.getOrgGUIDMutex.Lock()
	defer fake.getOrgGUIDMutex.Unlock()
	fake.GetOrgGUIDStub = nil
	fake.getOrgGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getOrgGUIDMutex.Lock()
	defer fake.getOrgGUIDMutex.Unlock()
	fake.GetOrgGUIDStub = nil
	if fake.getOrgGUIDReturnsOnCall == nil {
		fake.getOrgGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getOrgGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) ListOrgs() ([]cfclient.Org, error) {
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListOrgs", []interface{}{})
	fake.listOrgsMutex.Unlock()
	if fake.ListOrgsStub != nil {
		return fake.ListOrgsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listOrgsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeReader) ListOrgsCalls(stub func() ([]cfclient.Org, error)) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = stub
}

func (fake *FakeReader) ListOrgsReturns(result1 []cfclient.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) ListOrgsReturnsOnCall(i int, result1 []cfclient.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Org
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addOrgToListMutex.RLock()
	defer fake.addOrgToListMutex.RUnlock()
	fake.clearOrgListMutex.RLock()
	defer fake.clearOrgListMutex.RUnlock()
	fake.findOrgMutex.RLock()
	defer fake.findOrgMutex.RUnlock()
	fake.findOrgByGUIDMutex.RLock()
	defer fake.findOrgByGUIDMutex.RUnlock()
	fake.getOrgByGUIDMutex.RLock()
	defer fake.getOrgByGUIDMutex.RUnlock()
	fake.getOrgGUIDMutex.RLock()
	defer fake.getOrgGUIDMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReader) recordInvocation(key string, args []interface{}) {
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

var _ organizationreader.Reader = new(FakeReader)
