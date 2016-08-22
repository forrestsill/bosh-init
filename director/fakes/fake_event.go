// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/cloudfoundry/bosh-init/director"
)

type FakeEvent struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	ParentIDStub        func() string
	parentIDMutex       sync.RWMutex
	parentIDArgsForCall []struct{}
	parentIDReturns     struct {
		result1 string
	}
	TimestampStub        func() time.Time
	timestampMutex       sync.RWMutex
	timestampArgsForCall []struct{}
	timestampReturns     struct {
		result1 time.Time
	}
	UserStub        func() string
	userMutex       sync.RWMutex
	userArgsForCall []struct{}
	userReturns     struct {
		result1 string
	}
	ActionStub        func() string
	actionMutex       sync.RWMutex
	actionArgsForCall []struct{}
	actionReturns     struct {
		result1 string
	}
	ObjectTypeStub        func() string
	objectTypeMutex       sync.RWMutex
	objectTypeArgsForCall []struct{}
	objectTypeReturns     struct {
		result1 string
	}
	ObjectNameStub        func() string
	objectNameMutex       sync.RWMutex
	objectNameArgsForCall []struct{}
	objectNameReturns     struct {
		result1 string
	}
	TaskIDStub        func() string
	taskIDMutex       sync.RWMutex
	taskIDArgsForCall []struct{}
	taskIDReturns     struct {
		result1 string
	}
	DeploymentNameStub        func() string
	deploymentNameMutex       sync.RWMutex
	deploymentNameArgsForCall []struct{}
	deploymentNameReturns     struct {
		result1 string
	}
	InstanceStub        func() string
	instanceMutex       sync.RWMutex
	instanceArgsForCall []struct{}
	instanceReturns     struct {
		result1 string
	}
	ContextStub        func() map[string]interface{}
	contextMutex       sync.RWMutex
	contextArgsForCall []struct{}
	contextReturns     struct {
		result1 map[string]interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEvent) ID() string {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeEvent) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeEvent) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) ParentID() string {
	fake.parentIDMutex.Lock()
	fake.parentIDArgsForCall = append(fake.parentIDArgsForCall, struct{}{})
	fake.parentIDMutex.Unlock()
	if fake.ParentIDStub != nil {
		return fake.ParentIDStub()
	} else {
		return fake.parentIDReturns.result1
	}
}

func (fake *FakeEvent) ParentIDCallCount() int {
	fake.parentIDMutex.RLock()
	defer fake.parentIDMutex.RUnlock()
	return len(fake.parentIDArgsForCall)
}

func (fake *FakeEvent) ParentIDReturns(result1 string) {
	fake.ParentIDStub = nil
	fake.parentIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) Timestamp() time.Time {
	fake.timestampMutex.Lock()
	fake.timestampArgsForCall = append(fake.timestampArgsForCall, struct{}{})
	fake.timestampMutex.Unlock()
	if fake.TimestampStub != nil {
		return fake.TimestampStub()
	} else {
		return fake.timestampReturns.result1
	}
}

func (fake *FakeEvent) TimestampCallCount() int {
	fake.timestampMutex.RLock()
	defer fake.timestampMutex.RUnlock()
	return len(fake.timestampArgsForCall)
}

func (fake *FakeEvent) TimestampReturns(result1 time.Time) {
	fake.TimestampStub = nil
	fake.timestampReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeEvent) User() string {
	fake.userMutex.Lock()
	fake.userArgsForCall = append(fake.userArgsForCall, struct{}{})
	fake.userMutex.Unlock()
	if fake.UserStub != nil {
		return fake.UserStub()
	} else {
		return fake.userReturns.result1
	}
}

func (fake *FakeEvent) UserCallCount() int {
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	return len(fake.userArgsForCall)
}

func (fake *FakeEvent) UserReturns(result1 string) {
	fake.UserStub = nil
	fake.userReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) Action() string {
	fake.actionMutex.Lock()
	fake.actionArgsForCall = append(fake.actionArgsForCall, struct{}{})
	fake.actionMutex.Unlock()
	if fake.ActionStub != nil {
		return fake.ActionStub()
	} else {
		return fake.actionReturns.result1
	}
}

func (fake *FakeEvent) ActionCallCount() int {
	fake.actionMutex.RLock()
	defer fake.actionMutex.RUnlock()
	return len(fake.actionArgsForCall)
}

func (fake *FakeEvent) ActionReturns(result1 string) {
	fake.ActionStub = nil
	fake.actionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) ObjectType() string {
	fake.objectTypeMutex.Lock()
	fake.objectTypeArgsForCall = append(fake.objectTypeArgsForCall, struct{}{})
	fake.objectTypeMutex.Unlock()
	if fake.ObjectTypeStub != nil {
		return fake.ObjectTypeStub()
	} else {
		return fake.objectTypeReturns.result1
	}
}

func (fake *FakeEvent) ObjectTypeCallCount() int {
	fake.objectTypeMutex.RLock()
	defer fake.objectTypeMutex.RUnlock()
	return len(fake.objectTypeArgsForCall)
}

func (fake *FakeEvent) ObjectTypeReturns(result1 string) {
	fake.ObjectTypeStub = nil
	fake.objectTypeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) ObjectName() string {
	fake.objectNameMutex.Lock()
	fake.objectNameArgsForCall = append(fake.objectNameArgsForCall, struct{}{})
	fake.objectNameMutex.Unlock()
	if fake.ObjectNameStub != nil {
		return fake.ObjectNameStub()
	} else {
		return fake.objectNameReturns.result1
	}
}

func (fake *FakeEvent) ObjectNameCallCount() int {
	fake.objectNameMutex.RLock()
	defer fake.objectNameMutex.RUnlock()
	return len(fake.objectNameArgsForCall)
}

func (fake *FakeEvent) ObjectNameReturns(result1 string) {
	fake.ObjectNameStub = nil
	fake.objectNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) TaskID() string {
	fake.taskIDMutex.Lock()
	fake.taskIDArgsForCall = append(fake.taskIDArgsForCall, struct{}{})
	fake.taskIDMutex.Unlock()
	if fake.TaskIDStub != nil {
		return fake.TaskIDStub()
	} else {
		return fake.taskIDReturns.result1
	}
}

func (fake *FakeEvent) TaskIDCallCount() int {
	fake.taskIDMutex.RLock()
	defer fake.taskIDMutex.RUnlock()
	return len(fake.taskIDArgsForCall)
}

func (fake *FakeEvent) TaskIDReturns(result1 string) {
	fake.TaskIDStub = nil
	fake.taskIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) DeploymentName() string {
	fake.deploymentNameMutex.Lock()
	fake.deploymentNameArgsForCall = append(fake.deploymentNameArgsForCall, struct{}{})
	fake.deploymentNameMutex.Unlock()
	if fake.DeploymentNameStub != nil {
		return fake.DeploymentNameStub()
	} else {
		return fake.deploymentNameReturns.result1
	}
}

func (fake *FakeEvent) DeploymentNameCallCount() int {
	fake.deploymentNameMutex.RLock()
	defer fake.deploymentNameMutex.RUnlock()
	return len(fake.deploymentNameArgsForCall)
}

func (fake *FakeEvent) DeploymentNameReturns(result1 string) {
	fake.DeploymentNameStub = nil
	fake.deploymentNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) Instance() string {
	fake.instanceMutex.Lock()
	fake.instanceArgsForCall = append(fake.instanceArgsForCall, struct{}{})
	fake.instanceMutex.Unlock()
	if fake.InstanceStub != nil {
		return fake.InstanceStub()
	} else {
		return fake.instanceReturns.result1
	}
}

func (fake *FakeEvent) InstanceCallCount() int {
	fake.instanceMutex.RLock()
	defer fake.instanceMutex.RUnlock()
	return len(fake.instanceArgsForCall)
}

func (fake *FakeEvent) InstanceReturns(result1 string) {
	fake.InstanceStub = nil
	fake.instanceReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEvent) Context() map[string]interface{} {
	fake.contextMutex.Lock()
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	} else {
		return fake.contextReturns.result1
	}
}

func (fake *FakeEvent) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *FakeEvent) ContextReturns(result1 map[string]interface{}) {
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 map[string]interface{}
	}{result1}
}

var _ director.Event = new(FakeEvent)
