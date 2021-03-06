// This file was generated by counterfeiter
package directorfakes

import (
	"sync"
	"time"

	"github.com/cloudfoundry/bosh-init/director"
)

type FakeTask struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	CreatedAtStub        func() time.Time
	createdAtMutex       sync.RWMutex
	createdAtArgsForCall []struct{}
	createdAtReturns     struct {
		result1 time.Time
	}
	StateStub        func() string
	stateMutex       sync.RWMutex
	stateArgsForCall []struct{}
	stateReturns     struct {
		result1 string
	}
	IsErrorStub        func() bool
	isErrorMutex       sync.RWMutex
	isErrorArgsForCall []struct{}
	isErrorReturns     struct {
		result1 bool
	}
	UserStub        func() string
	userMutex       sync.RWMutex
	userArgsForCall []struct{}
	userReturns     struct {
		result1 string
	}
	DescriptionStub        func() string
	descriptionMutex       sync.RWMutex
	descriptionArgsForCall []struct{}
	descriptionReturns     struct {
		result1 string
	}
	ResultStub        func() string
	resultMutex       sync.RWMutex
	resultArgsForCall []struct{}
	resultReturns     struct {
		result1 string
	}
	EventOutputStub        func(director.TaskReporter) error
	eventOutputMutex       sync.RWMutex
	eventOutputArgsForCall []struct {
		arg1 director.TaskReporter
	}
	eventOutputReturns struct {
		result1 error
	}
	CPIOutputStub        func(director.TaskReporter) error
	cPIOutputMutex       sync.RWMutex
	cPIOutputArgsForCall []struct {
		arg1 director.TaskReporter
	}
	cPIOutputReturns struct {
		result1 error
	}
	DebugOutputStub        func(director.TaskReporter) error
	debugOutputMutex       sync.RWMutex
	debugOutputArgsForCall []struct {
		arg1 director.TaskReporter
	}
	debugOutputReturns struct {
		result1 error
	}
	ResultOutputStub        func(director.TaskReporter) error
	resultOutputMutex       sync.RWMutex
	resultOutputArgsForCall []struct {
		arg1 director.TaskReporter
	}
	resultOutputReturns struct {
		result1 error
	}
	RawOutputStub        func(director.TaskReporter) error
	rawOutputMutex       sync.RWMutex
	rawOutputArgsForCall []struct {
		arg1 director.TaskReporter
	}
	rawOutputReturns struct {
		result1 error
	}
	CancelStub        func() error
	cancelMutex       sync.RWMutex
	cancelArgsForCall []struct{}
	cancelReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTask) ID() int {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeTask) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeTask) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeTask) CreatedAt() time.Time {
	fake.createdAtMutex.Lock()
	fake.createdAtArgsForCall = append(fake.createdAtArgsForCall, struct{}{})
	fake.recordInvocation("CreatedAt", []interface{}{})
	fake.createdAtMutex.Unlock()
	if fake.CreatedAtStub != nil {
		return fake.CreatedAtStub()
	} else {
		return fake.createdAtReturns.result1
	}
}

func (fake *FakeTask) CreatedAtCallCount() int {
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	return len(fake.createdAtArgsForCall)
}

func (fake *FakeTask) CreatedAtReturns(result1 time.Time) {
	fake.CreatedAtStub = nil
	fake.createdAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeTask) State() string {
	fake.stateMutex.Lock()
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct{}{})
	fake.recordInvocation("State", []interface{}{})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub()
	} else {
		return fake.stateReturns.result1
	}
}

func (fake *FakeTask) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeTask) StateReturns(result1 string) {
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) IsError() bool {
	fake.isErrorMutex.Lock()
	fake.isErrorArgsForCall = append(fake.isErrorArgsForCall, struct{}{})
	fake.recordInvocation("IsError", []interface{}{})
	fake.isErrorMutex.Unlock()
	if fake.IsErrorStub != nil {
		return fake.IsErrorStub()
	} else {
		return fake.isErrorReturns.result1
	}
}

func (fake *FakeTask) IsErrorCallCount() int {
	fake.isErrorMutex.RLock()
	defer fake.isErrorMutex.RUnlock()
	return len(fake.isErrorArgsForCall)
}

func (fake *FakeTask) IsErrorReturns(result1 bool) {
	fake.IsErrorStub = nil
	fake.isErrorReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeTask) User() string {
	fake.userMutex.Lock()
	fake.userArgsForCall = append(fake.userArgsForCall, struct{}{})
	fake.recordInvocation("User", []interface{}{})
	fake.userMutex.Unlock()
	if fake.UserStub != nil {
		return fake.UserStub()
	} else {
		return fake.userReturns.result1
	}
}

func (fake *FakeTask) UserCallCount() int {
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	return len(fake.userArgsForCall)
}

func (fake *FakeTask) UserReturns(result1 string) {
	fake.UserStub = nil
	fake.userReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) Description() string {
	fake.descriptionMutex.Lock()
	fake.descriptionArgsForCall = append(fake.descriptionArgsForCall, struct{}{})
	fake.recordInvocation("Description", []interface{}{})
	fake.descriptionMutex.Unlock()
	if fake.DescriptionStub != nil {
		return fake.DescriptionStub()
	} else {
		return fake.descriptionReturns.result1
	}
}

func (fake *FakeTask) DescriptionCallCount() int {
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	return len(fake.descriptionArgsForCall)
}

func (fake *FakeTask) DescriptionReturns(result1 string) {
	fake.DescriptionStub = nil
	fake.descriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) Result() string {
	fake.resultMutex.Lock()
	fake.resultArgsForCall = append(fake.resultArgsForCall, struct{}{})
	fake.recordInvocation("Result", []interface{}{})
	fake.resultMutex.Unlock()
	if fake.ResultStub != nil {
		return fake.ResultStub()
	} else {
		return fake.resultReturns.result1
	}
}

func (fake *FakeTask) ResultCallCount() int {
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	return len(fake.resultArgsForCall)
}

func (fake *FakeTask) ResultReturns(result1 string) {
	fake.ResultStub = nil
	fake.resultReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) EventOutput(arg1 director.TaskReporter) error {
	fake.eventOutputMutex.Lock()
	fake.eventOutputArgsForCall = append(fake.eventOutputArgsForCall, struct {
		arg1 director.TaskReporter
	}{arg1})
	fake.recordInvocation("EventOutput", []interface{}{arg1})
	fake.eventOutputMutex.Unlock()
	if fake.EventOutputStub != nil {
		return fake.EventOutputStub(arg1)
	} else {
		return fake.eventOutputReturns.result1
	}
}

func (fake *FakeTask) EventOutputCallCount() int {
	fake.eventOutputMutex.RLock()
	defer fake.eventOutputMutex.RUnlock()
	return len(fake.eventOutputArgsForCall)
}

func (fake *FakeTask) EventOutputArgsForCall(i int) director.TaskReporter {
	fake.eventOutputMutex.RLock()
	defer fake.eventOutputMutex.RUnlock()
	return fake.eventOutputArgsForCall[i].arg1
}

func (fake *FakeTask) EventOutputReturns(result1 error) {
	fake.EventOutputStub = nil
	fake.eventOutputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) CPIOutput(arg1 director.TaskReporter) error {
	fake.cPIOutputMutex.Lock()
	fake.cPIOutputArgsForCall = append(fake.cPIOutputArgsForCall, struct {
		arg1 director.TaskReporter
	}{arg1})
	fake.recordInvocation("CPIOutput", []interface{}{arg1})
	fake.cPIOutputMutex.Unlock()
	if fake.CPIOutputStub != nil {
		return fake.CPIOutputStub(arg1)
	} else {
		return fake.cPIOutputReturns.result1
	}
}

func (fake *FakeTask) CPIOutputCallCount() int {
	fake.cPIOutputMutex.RLock()
	defer fake.cPIOutputMutex.RUnlock()
	return len(fake.cPIOutputArgsForCall)
}

func (fake *FakeTask) CPIOutputArgsForCall(i int) director.TaskReporter {
	fake.cPIOutputMutex.RLock()
	defer fake.cPIOutputMutex.RUnlock()
	return fake.cPIOutputArgsForCall[i].arg1
}

func (fake *FakeTask) CPIOutputReturns(result1 error) {
	fake.CPIOutputStub = nil
	fake.cPIOutputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) DebugOutput(arg1 director.TaskReporter) error {
	fake.debugOutputMutex.Lock()
	fake.debugOutputArgsForCall = append(fake.debugOutputArgsForCall, struct {
		arg1 director.TaskReporter
	}{arg1})
	fake.recordInvocation("DebugOutput", []interface{}{arg1})
	fake.debugOutputMutex.Unlock()
	if fake.DebugOutputStub != nil {
		return fake.DebugOutputStub(arg1)
	} else {
		return fake.debugOutputReturns.result1
	}
}

func (fake *FakeTask) DebugOutputCallCount() int {
	fake.debugOutputMutex.RLock()
	defer fake.debugOutputMutex.RUnlock()
	return len(fake.debugOutputArgsForCall)
}

func (fake *FakeTask) DebugOutputArgsForCall(i int) director.TaskReporter {
	fake.debugOutputMutex.RLock()
	defer fake.debugOutputMutex.RUnlock()
	return fake.debugOutputArgsForCall[i].arg1
}

func (fake *FakeTask) DebugOutputReturns(result1 error) {
	fake.DebugOutputStub = nil
	fake.debugOutputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) ResultOutput(arg1 director.TaskReporter) error {
	fake.resultOutputMutex.Lock()
	fake.resultOutputArgsForCall = append(fake.resultOutputArgsForCall, struct {
		arg1 director.TaskReporter
	}{arg1})
	fake.recordInvocation("ResultOutput", []interface{}{arg1})
	fake.resultOutputMutex.Unlock()
	if fake.ResultOutputStub != nil {
		return fake.ResultOutputStub(arg1)
	} else {
		return fake.resultOutputReturns.result1
	}
}

func (fake *FakeTask) ResultOutputCallCount() int {
	fake.resultOutputMutex.RLock()
	defer fake.resultOutputMutex.RUnlock()
	return len(fake.resultOutputArgsForCall)
}

func (fake *FakeTask) ResultOutputArgsForCall(i int) director.TaskReporter {
	fake.resultOutputMutex.RLock()
	defer fake.resultOutputMutex.RUnlock()
	return fake.resultOutputArgsForCall[i].arg1
}

func (fake *FakeTask) ResultOutputReturns(result1 error) {
	fake.ResultOutputStub = nil
	fake.resultOutputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) RawOutput(arg1 director.TaskReporter) error {
	fake.rawOutputMutex.Lock()
	fake.rawOutputArgsForCall = append(fake.rawOutputArgsForCall, struct {
		arg1 director.TaskReporter
	}{arg1})
	fake.recordInvocation("RawOutput", []interface{}{arg1})
	fake.rawOutputMutex.Unlock()
	if fake.RawOutputStub != nil {
		return fake.RawOutputStub(arg1)
	} else {
		return fake.rawOutputReturns.result1
	}
}

func (fake *FakeTask) RawOutputCallCount() int {
	fake.rawOutputMutex.RLock()
	defer fake.rawOutputMutex.RUnlock()
	return len(fake.rawOutputArgsForCall)
}

func (fake *FakeTask) RawOutputArgsForCall(i int) director.TaskReporter {
	fake.rawOutputMutex.RLock()
	defer fake.rawOutputMutex.RUnlock()
	return fake.rawOutputArgsForCall[i].arg1
}

func (fake *FakeTask) RawOutputReturns(result1 error) {
	fake.RawOutputStub = nil
	fake.rawOutputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) Cancel() error {
	fake.cancelMutex.Lock()
	fake.cancelArgsForCall = append(fake.cancelArgsForCall, struct{}{})
	fake.recordInvocation("Cancel", []interface{}{})
	fake.cancelMutex.Unlock()
	if fake.CancelStub != nil {
		return fake.CancelStub()
	} else {
		return fake.cancelReturns.result1
	}
}

func (fake *FakeTask) CancelCallCount() int {
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	return len(fake.cancelArgsForCall)
}

func (fake *FakeTask) CancelReturns(result1 error) {
	fake.CancelStub = nil
	fake.cancelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.isErrorMutex.RLock()
	defer fake.isErrorMutex.RUnlock()
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	fake.eventOutputMutex.RLock()
	defer fake.eventOutputMutex.RUnlock()
	fake.cPIOutputMutex.RLock()
	defer fake.cPIOutputMutex.RUnlock()
	fake.debugOutputMutex.RLock()
	defer fake.debugOutputMutex.RUnlock()
	fake.resultOutputMutex.RLock()
	defer fake.resultOutputMutex.RUnlock()
	fake.rawOutputMutex.RLock()
	defer fake.rawOutputMutex.RUnlock()
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTask) recordInvocation(key string, args []interface{}) {
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

var _ director.Task = new(FakeTask)
