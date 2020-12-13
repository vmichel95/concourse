// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/worker"
)

type FakeArtifactWirer struct {
	WireImageStub        func(lager.Logger, runtime.Artifact) (worker.StreamableArtifactSource, error)
	wireImageMutex       sync.RWMutex
	wireImageArgsForCall []struct {
		arg1 lager.Logger
		arg2 runtime.Artifact
	}
	wireImageReturns struct {
		result1 worker.StreamableArtifactSource
		result2 error
	}
	wireImageReturnsOnCall map[int]struct {
		result1 worker.StreamableArtifactSource
		result2 error
	}
	WireInputsAndCachesStub        func(lager.Logger, int, map[string]runtime.Artifact) ([]worker.InputSource, error)
	wireInputsAndCachesMutex       sync.RWMutex
	wireInputsAndCachesArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 map[string]runtime.Artifact
	}
	wireInputsAndCachesReturns struct {
		result1 []worker.InputSource
		result2 error
	}
	wireInputsAndCachesReturnsOnCall map[int]struct {
		result1 []worker.InputSource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArtifactWirer) WireImage(arg1 lager.Logger, arg2 runtime.Artifact) (worker.StreamableArtifactSource, error) {
	fake.wireImageMutex.Lock()
	ret, specificReturn := fake.wireImageReturnsOnCall[len(fake.wireImageArgsForCall)]
	fake.wireImageArgsForCall = append(fake.wireImageArgsForCall, struct {
		arg1 lager.Logger
		arg2 runtime.Artifact
	}{arg1, arg2})
	fake.recordInvocation("WireImage", []interface{}{arg1, arg2})
	fake.wireImageMutex.Unlock()
	if fake.WireImageStub != nil {
		return fake.WireImageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.wireImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeArtifactWirer) WireImageCallCount() int {
	fake.wireImageMutex.RLock()
	defer fake.wireImageMutex.RUnlock()
	return len(fake.wireImageArgsForCall)
}

func (fake *FakeArtifactWirer) WireImageCalls(stub func(lager.Logger, runtime.Artifact) (worker.StreamableArtifactSource, error)) {
	fake.wireImageMutex.Lock()
	defer fake.wireImageMutex.Unlock()
	fake.WireImageStub = stub
}

func (fake *FakeArtifactWirer) WireImageArgsForCall(i int) (lager.Logger, runtime.Artifact) {
	fake.wireImageMutex.RLock()
	defer fake.wireImageMutex.RUnlock()
	argsForCall := fake.wireImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeArtifactWirer) WireImageReturns(result1 worker.StreamableArtifactSource, result2 error) {
	fake.wireImageMutex.Lock()
	defer fake.wireImageMutex.Unlock()
	fake.WireImageStub = nil
	fake.wireImageReturns = struct {
		result1 worker.StreamableArtifactSource
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifactWirer) WireImageReturnsOnCall(i int, result1 worker.StreamableArtifactSource, result2 error) {
	fake.wireImageMutex.Lock()
	defer fake.wireImageMutex.Unlock()
	fake.WireImageStub = nil
	if fake.wireImageReturnsOnCall == nil {
		fake.wireImageReturnsOnCall = make(map[int]struct {
			result1 worker.StreamableArtifactSource
			result2 error
		})
	}
	fake.wireImageReturnsOnCall[i] = struct {
		result1 worker.StreamableArtifactSource
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifactWirer) WireInputsAndCaches(arg1 lager.Logger, arg2 int, arg3 map[string]runtime.Artifact) ([]worker.InputSource, error) {
	fake.wireInputsAndCachesMutex.Lock()
	ret, specificReturn := fake.wireInputsAndCachesReturnsOnCall[len(fake.wireInputsAndCachesArgsForCall)]
	fake.wireInputsAndCachesArgsForCall = append(fake.wireInputsAndCachesArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 map[string]runtime.Artifact
	}{arg1, arg2, arg3})
	fake.recordInvocation("WireInputsAndCaches", []interface{}{arg1, arg2, arg3})
	fake.wireInputsAndCachesMutex.Unlock()
	if fake.WireInputsAndCachesStub != nil {
		return fake.WireInputsAndCachesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.wireInputsAndCachesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeArtifactWirer) WireInputsAndCachesCallCount() int {
	fake.wireInputsAndCachesMutex.RLock()
	defer fake.wireInputsAndCachesMutex.RUnlock()
	return len(fake.wireInputsAndCachesArgsForCall)
}

func (fake *FakeArtifactWirer) WireInputsAndCachesCalls(stub func(lager.Logger, int, map[string]runtime.Artifact) ([]worker.InputSource, error)) {
	fake.wireInputsAndCachesMutex.Lock()
	defer fake.wireInputsAndCachesMutex.Unlock()
	fake.WireInputsAndCachesStub = stub
}

func (fake *FakeArtifactWirer) WireInputsAndCachesArgsForCall(i int) (lager.Logger, int, map[string]runtime.Artifact) {
	fake.wireInputsAndCachesMutex.RLock()
	defer fake.wireInputsAndCachesMutex.RUnlock()
	argsForCall := fake.wireInputsAndCachesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeArtifactWirer) WireInputsAndCachesReturns(result1 []worker.InputSource, result2 error) {
	fake.wireInputsAndCachesMutex.Lock()
	defer fake.wireInputsAndCachesMutex.Unlock()
	fake.WireInputsAndCachesStub = nil
	fake.wireInputsAndCachesReturns = struct {
		result1 []worker.InputSource
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifactWirer) WireInputsAndCachesReturnsOnCall(i int, result1 []worker.InputSource, result2 error) {
	fake.wireInputsAndCachesMutex.Lock()
	defer fake.wireInputsAndCachesMutex.Unlock()
	fake.WireInputsAndCachesStub = nil
	if fake.wireInputsAndCachesReturnsOnCall == nil {
		fake.wireInputsAndCachesReturnsOnCall = make(map[int]struct {
			result1 []worker.InputSource
			result2 error
		})
	}
	fake.wireInputsAndCachesReturnsOnCall[i] = struct {
		result1 []worker.InputSource
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifactWirer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.wireImageMutex.RLock()
	defer fake.wireImageMutex.RUnlock()
	fake.wireInputsAndCachesMutex.RLock()
	defer fake.wireInputsAndCachesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArtifactWirer) recordInvocation(key string, args []interface{}) {
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

var _ worker.ArtifactWirer = new(FakeArtifactWirer)