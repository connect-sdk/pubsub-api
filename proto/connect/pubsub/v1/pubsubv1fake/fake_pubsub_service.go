// Code generated by counterfeiter. DO NOT EDIT.
package pubsubv1fake

import (
	"context"
	"sync"

	pubsubv1 "github.com/connect-sdk/pubsub-api/proto/connect/pubsub/v1"
)

type FakePubsubService struct {
	PushPubsubMessageStub        func(context.Context, *pubsubv1.PushPubsubMessageRequest) (*pubsubv1.PushPubsubMessageResponse, error)
	pushPubsubMessageMutex       sync.RWMutex
	pushPubsubMessageArgsForCall []struct {
		arg1 context.Context
		arg2 *pubsubv1.PushPubsubMessageRequest
	}
	pushPubsubMessageReturns struct {
		result1 *pubsubv1.PushPubsubMessageResponse
		result2 error
	}
	pushPubsubMessageReturnsOnCall map[int]struct {
		result1 *pubsubv1.PushPubsubMessageResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePubsubService) PushPubsubMessage(arg1 context.Context, arg2 *pubsubv1.PushPubsubMessageRequest) (*pubsubv1.PushPubsubMessageResponse, error) {
	fake.pushPubsubMessageMutex.Lock()
	ret, specificReturn := fake.pushPubsubMessageReturnsOnCall[len(fake.pushPubsubMessageArgsForCall)]
	fake.pushPubsubMessageArgsForCall = append(fake.pushPubsubMessageArgsForCall, struct {
		arg1 context.Context
		arg2 *pubsubv1.PushPubsubMessageRequest
	}{arg1, arg2})
	stub := fake.PushPubsubMessageStub
	fakeReturns := fake.pushPubsubMessageReturns
	fake.recordInvocation("PushPubsubMessage", []interface{}{arg1, arg2})
	fake.pushPubsubMessageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePubsubService) PushPubsubMessageCallCount() int {
	fake.pushPubsubMessageMutex.RLock()
	defer fake.pushPubsubMessageMutex.RUnlock()
	return len(fake.pushPubsubMessageArgsForCall)
}

func (fake *FakePubsubService) PushPubsubMessageCalls(stub func(context.Context, *pubsubv1.PushPubsubMessageRequest) (*pubsubv1.PushPubsubMessageResponse, error)) {
	fake.pushPubsubMessageMutex.Lock()
	defer fake.pushPubsubMessageMutex.Unlock()
	fake.PushPubsubMessageStub = stub
}

func (fake *FakePubsubService) PushPubsubMessageArgsForCall(i int) (context.Context, *pubsubv1.PushPubsubMessageRequest) {
	fake.pushPubsubMessageMutex.RLock()
	defer fake.pushPubsubMessageMutex.RUnlock()
	argsForCall := fake.pushPubsubMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePubsubService) PushPubsubMessageReturns(result1 *pubsubv1.PushPubsubMessageResponse, result2 error) {
	fake.pushPubsubMessageMutex.Lock()
	defer fake.pushPubsubMessageMutex.Unlock()
	fake.PushPubsubMessageStub = nil
	fake.pushPubsubMessageReturns = struct {
		result1 *pubsubv1.PushPubsubMessageResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePubsubService) PushPubsubMessageReturnsOnCall(i int, result1 *pubsubv1.PushPubsubMessageResponse, result2 error) {
	fake.pushPubsubMessageMutex.Lock()
	defer fake.pushPubsubMessageMutex.Unlock()
	fake.PushPubsubMessageStub = nil
	if fake.pushPubsubMessageReturnsOnCall == nil {
		fake.pushPubsubMessageReturnsOnCall = make(map[int]struct {
			result1 *pubsubv1.PushPubsubMessageResponse
			result2 error
		})
	}
	fake.pushPubsubMessageReturnsOnCall[i] = struct {
		result1 *pubsubv1.PushPubsubMessageResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePubsubService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pushPubsubMessageMutex.RLock()
	defer fake.pushPubsubMessageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePubsubService) recordInvocation(key string, args []interface{}) {
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

var _ pubsubv1.PubsubService = new(FakePubsubService)
