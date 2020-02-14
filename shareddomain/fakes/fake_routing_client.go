// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/routing-api/models"
	"github.com/pivotal-michael-stergianis/cf-mgmt/shareddomain"
)

type FakeRoutingClient struct {
	RouterGroupWithNameStub        func(string) (models.RouterGroup, error)
	routerGroupWithNameMutex       sync.RWMutex
	routerGroupWithNameArgsForCall []struct {
		arg1 string
	}
	routerGroupWithNameReturns struct {
		result1 models.RouterGroup
		result2 error
	}
	routerGroupWithNameReturnsOnCall map[int]struct {
		result1 models.RouterGroup
		result2 error
	}
	RouterGroupsStub        func() ([]models.RouterGroup, error)
	routerGroupsMutex       sync.RWMutex
	routerGroupsArgsForCall []struct {
	}
	routerGroupsReturns struct {
		result1 []models.RouterGroup
		result2 error
	}
	routerGroupsReturnsOnCall map[int]struct {
		result1 []models.RouterGroup
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoutingClient) RouterGroupWithName(arg1 string) (models.RouterGroup, error) {
	fake.routerGroupWithNameMutex.Lock()
	ret, specificReturn := fake.routerGroupWithNameReturnsOnCall[len(fake.routerGroupWithNameArgsForCall)]
	fake.routerGroupWithNameArgsForCall = append(fake.routerGroupWithNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RouterGroupWithName", []interface{}{arg1})
	fake.routerGroupWithNameMutex.Unlock()
	if fake.RouterGroupWithNameStub != nil {
		return fake.RouterGroupWithNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.routerGroupWithNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRoutingClient) RouterGroupWithNameCallCount() int {
	fake.routerGroupWithNameMutex.RLock()
	defer fake.routerGroupWithNameMutex.RUnlock()
	return len(fake.routerGroupWithNameArgsForCall)
}

func (fake *FakeRoutingClient) RouterGroupWithNameCalls(stub func(string) (models.RouterGroup, error)) {
	fake.routerGroupWithNameMutex.Lock()
	defer fake.routerGroupWithNameMutex.Unlock()
	fake.RouterGroupWithNameStub = stub
}

func (fake *FakeRoutingClient) RouterGroupWithNameArgsForCall(i int) string {
	fake.routerGroupWithNameMutex.RLock()
	defer fake.routerGroupWithNameMutex.RUnlock()
	argsForCall := fake.routerGroupWithNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRoutingClient) RouterGroupWithNameReturns(result1 models.RouterGroup, result2 error) {
	fake.routerGroupWithNameMutex.Lock()
	defer fake.routerGroupWithNameMutex.Unlock()
	fake.RouterGroupWithNameStub = nil
	fake.routerGroupWithNameReturns = struct {
		result1 models.RouterGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeRoutingClient) RouterGroupWithNameReturnsOnCall(i int, result1 models.RouterGroup, result2 error) {
	fake.routerGroupWithNameMutex.Lock()
	defer fake.routerGroupWithNameMutex.Unlock()
	fake.RouterGroupWithNameStub = nil
	if fake.routerGroupWithNameReturnsOnCall == nil {
		fake.routerGroupWithNameReturnsOnCall = make(map[int]struct {
			result1 models.RouterGroup
			result2 error
		})
	}
	fake.routerGroupWithNameReturnsOnCall[i] = struct {
		result1 models.RouterGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeRoutingClient) RouterGroups() ([]models.RouterGroup, error) {
	fake.routerGroupsMutex.Lock()
	ret, specificReturn := fake.routerGroupsReturnsOnCall[len(fake.routerGroupsArgsForCall)]
	fake.routerGroupsArgsForCall = append(fake.routerGroupsArgsForCall, struct {
	}{})
	fake.recordInvocation("RouterGroups", []interface{}{})
	fake.routerGroupsMutex.Unlock()
	if fake.RouterGroupsStub != nil {
		return fake.RouterGroupsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.routerGroupsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRoutingClient) RouterGroupsCallCount() int {
	fake.routerGroupsMutex.RLock()
	defer fake.routerGroupsMutex.RUnlock()
	return len(fake.routerGroupsArgsForCall)
}

func (fake *FakeRoutingClient) RouterGroupsCalls(stub func() ([]models.RouterGroup, error)) {
	fake.routerGroupsMutex.Lock()
	defer fake.routerGroupsMutex.Unlock()
	fake.RouterGroupsStub = stub
}

func (fake *FakeRoutingClient) RouterGroupsReturns(result1 []models.RouterGroup, result2 error) {
	fake.routerGroupsMutex.Lock()
	defer fake.routerGroupsMutex.Unlock()
	fake.RouterGroupsStub = nil
	fake.routerGroupsReturns = struct {
		result1 []models.RouterGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeRoutingClient) RouterGroupsReturnsOnCall(i int, result1 []models.RouterGroup, result2 error) {
	fake.routerGroupsMutex.Lock()
	defer fake.routerGroupsMutex.Unlock()
	fake.RouterGroupsStub = nil
	if fake.routerGroupsReturnsOnCall == nil {
		fake.routerGroupsReturnsOnCall = make(map[int]struct {
			result1 []models.RouterGroup
			result2 error
		})
	}
	fake.routerGroupsReturnsOnCall[i] = struct {
		result1 []models.RouterGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeRoutingClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.routerGroupWithNameMutex.RLock()
	defer fake.routerGroupWithNameMutex.RUnlock()
	fake.routerGroupsMutex.RLock()
	defer fake.routerGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoutingClient) recordInvocation(key string, args []interface{}) {
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

var _ shareddomain.RoutingClient = new(FakeRoutingClient)
