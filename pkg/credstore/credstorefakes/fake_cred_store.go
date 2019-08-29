// Code generated by counterfeiter. DO NOT EDIT.
package credstorefakes

import (
	"sync"

	"code.cloudfoundry.org/credhub-cli/credhub/permissions"
	"github.com/cf-platform-eng/kibosh/pkg/credstore"
)

type FakeCredStore struct {
	AddPermissionStub        func(string, string, []string) (*permissions.Permission, error)
	addPermissionMutex       sync.RWMutex
	addPermissionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []string
	}
	addPermissionReturns struct {
		result1 *permissions.Permission
		result2 error
	}
	addPermissionReturnsOnCall map[int]struct {
		result1 *permissions.Permission
		result2 error
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeletePermissionStub        func(string) error
	deletePermissionMutex       sync.RWMutex
	deletePermissionArgsForCall []struct {
		arg1 string
	}
	deletePermissionReturns struct {
		result1 error
	}
	deletePermissionReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string) (string, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 string
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	PutStub        func(string, interface{}) (interface{}, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	putReturns struct {
		result1 interface{}
		result2 error
	}
	putReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredStore) AddPermission(arg1 string, arg2 string, arg3 []string) (*permissions.Permission, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.addPermissionMutex.Lock()
	ret, specificReturn := fake.addPermissionReturnsOnCall[len(fake.addPermissionArgsForCall)]
	fake.addPermissionArgsForCall = append(fake.addPermissionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("AddPermission", []interface{}{arg1, arg2, arg3Copy})
	fake.addPermissionMutex.Unlock()
	if fake.AddPermissionStub != nil {
		return fake.AddPermissionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.addPermissionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredStore) AddPermissionCallCount() int {
	fake.addPermissionMutex.RLock()
	defer fake.addPermissionMutex.RUnlock()
	return len(fake.addPermissionArgsForCall)
}

func (fake *FakeCredStore) AddPermissionCalls(stub func(string, string, []string) (*permissions.Permission, error)) {
	fake.addPermissionMutex.Lock()
	defer fake.addPermissionMutex.Unlock()
	fake.AddPermissionStub = stub
}

func (fake *FakeCredStore) AddPermissionArgsForCall(i int) (string, string, []string) {
	fake.addPermissionMutex.RLock()
	defer fake.addPermissionMutex.RUnlock()
	argsForCall := fake.addPermissionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCredStore) AddPermissionReturns(result1 *permissions.Permission, result2 error) {
	fake.addPermissionMutex.Lock()
	defer fake.addPermissionMutex.Unlock()
	fake.AddPermissionStub = nil
	fake.addPermissionReturns = struct {
		result1 *permissions.Permission
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) AddPermissionReturnsOnCall(i int, result1 *permissions.Permission, result2 error) {
	fake.addPermissionMutex.Lock()
	defer fake.addPermissionMutex.Unlock()
	fake.AddPermissionStub = nil
	if fake.addPermissionReturnsOnCall == nil {
		fake.addPermissionReturnsOnCall = make(map[int]struct {
			result1 *permissions.Permission
			result2 error
		})
	}
	fake.addPermissionReturnsOnCall[i] = struct {
		result1 *permissions.Permission
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeCredStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCredStore) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeCredStore) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredStore) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredStore) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredStore) DeletePermission(arg1 string) error {
	fake.deletePermissionMutex.Lock()
	ret, specificReturn := fake.deletePermissionReturnsOnCall[len(fake.deletePermissionArgsForCall)]
	fake.deletePermissionArgsForCall = append(fake.deletePermissionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeletePermission", []interface{}{arg1})
	fake.deletePermissionMutex.Unlock()
	if fake.DeletePermissionStub != nil {
		return fake.DeletePermissionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deletePermissionReturns
	return fakeReturns.result1
}

func (fake *FakeCredStore) DeletePermissionCallCount() int {
	fake.deletePermissionMutex.RLock()
	defer fake.deletePermissionMutex.RUnlock()
	return len(fake.deletePermissionArgsForCall)
}

func (fake *FakeCredStore) DeletePermissionCalls(stub func(string) error) {
	fake.deletePermissionMutex.Lock()
	defer fake.deletePermissionMutex.Unlock()
	fake.DeletePermissionStub = stub
}

func (fake *FakeCredStore) DeletePermissionArgsForCall(i int) string {
	fake.deletePermissionMutex.RLock()
	defer fake.deletePermissionMutex.RUnlock()
	argsForCall := fake.deletePermissionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredStore) DeletePermissionReturns(result1 error) {
	fake.deletePermissionMutex.Lock()
	defer fake.deletePermissionMutex.Unlock()
	fake.DeletePermissionStub = nil
	fake.deletePermissionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredStore) DeletePermissionReturnsOnCall(i int, result1 error) {
	fake.deletePermissionMutex.Lock()
	defer fake.deletePermissionMutex.Unlock()
	fake.DeletePermissionStub = nil
	if fake.deletePermissionReturnsOnCall == nil {
		fake.deletePermissionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deletePermissionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredStore) Get(arg1 string) (string, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCredStore) GetCalls(stub func(string) (string, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeCredStore) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredStore) GetReturns(result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) GetReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) Put(arg1 string, arg2 interface{}) (interface{}, error) {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("Put", []interface{}{arg1, arg2})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredStore) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeCredStore) PutCalls(stub func(string, interface{}) (interface{}, error)) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = stub
}

func (fake *FakeCredStore) PutArgsForCall(i int) (string, interface{}) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	argsForCall := fake.putArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCredStore) PutReturns(result1 interface{}, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) PutReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addPermissionMutex.RLock()
	defer fake.addPermissionMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deletePermissionMutex.RLock()
	defer fake.deletePermissionMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredStore) recordInvocation(key string, args []interface{}) {
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

var _ credstore.CredStore = new(FakeCredStore)