// Code generated by counterfeiter. DO NOT EDIT.
package grootfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/groot"
	"code.cloudfoundry.org/lager/v3"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeDriver struct {
	BundleStub        func(lager.Logger, string, []string, int64) (specs.Spec, error)
	bundleMutex       sync.RWMutex
	bundleArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 []string
		arg4 int64
	}
	bundleReturns struct {
		result1 specs.Spec
		result2 error
	}
	bundleReturnsOnCall map[int]struct {
		result1 specs.Spec
		result2 error
	}
	DeleteStub        func(lager.Logger, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	StatsStub        func(lager.Logger, string) (groot.VolumeStats, error)
	statsMutex       sync.RWMutex
	statsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	statsReturns struct {
		result1 groot.VolumeStats
		result2 error
	}
	statsReturnsOnCall map[int]struct {
		result1 groot.VolumeStats
		result2 error
	}
	UnpackStub        func(lager.Logger, string, []string, io.Reader) (int64, error)
	unpackMutex       sync.RWMutex
	unpackArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 []string
		arg4 io.Reader
	}
	unpackReturns struct {
		result1 int64
		result2 error
	}
	unpackReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	WriteMetadataStub        func(lager.Logger, string, groot.ImageMetadata) error
	writeMetadataMutex       sync.RWMutex
	writeMetadataArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 groot.ImageMetadata
	}
	writeMetadataReturns struct {
		result1 error
	}
	writeMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDriver) Bundle(arg1 lager.Logger, arg2 string, arg3 []string, arg4 int64) (specs.Spec, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.bundleMutex.Lock()
	ret, specificReturn := fake.bundleReturnsOnCall[len(fake.bundleArgsForCall)]
	fake.bundleArgsForCall = append(fake.bundleArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 []string
		arg4 int64
	}{arg1, arg2, arg3Copy, arg4})
	stub := fake.BundleStub
	fakeReturns := fake.bundleReturns
	fake.recordInvocation("Bundle", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.bundleMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDriver) BundleCallCount() int {
	fake.bundleMutex.RLock()
	defer fake.bundleMutex.RUnlock()
	return len(fake.bundleArgsForCall)
}

func (fake *FakeDriver) BundleCalls(stub func(lager.Logger, string, []string, int64) (specs.Spec, error)) {
	fake.bundleMutex.Lock()
	defer fake.bundleMutex.Unlock()
	fake.BundleStub = stub
}

func (fake *FakeDriver) BundleArgsForCall(i int) (lager.Logger, string, []string, int64) {
	fake.bundleMutex.RLock()
	defer fake.bundleMutex.RUnlock()
	argsForCall := fake.bundleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDriver) BundleReturns(result1 specs.Spec, result2 error) {
	fake.bundleMutex.Lock()
	defer fake.bundleMutex.Unlock()
	fake.BundleStub = nil
	fake.bundleReturns = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeDriver) BundleReturnsOnCall(i int, result1 specs.Spec, result2 error) {
	fake.bundleMutex.Lock()
	defer fake.bundleMutex.Unlock()
	fake.BundleStub = nil
	if fake.bundleReturnsOnCall == nil {
		fake.bundleReturnsOnCall = make(map[int]struct {
			result1 specs.Spec
			result2 error
		})
	}
	fake.bundleReturnsOnCall[i] = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeDriver) Delete(arg1 lager.Logger, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDriver) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeDriver) DeleteCalls(stub func(lager.Logger, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeDriver) DeleteArgsForCall(i int) (lager.Logger, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDriver) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeDriver) Stats(arg1 lager.Logger, arg2 string) (groot.VolumeStats, error) {
	fake.statsMutex.Lock()
	ret, specificReturn := fake.statsReturnsOnCall[len(fake.statsArgsForCall)]
	fake.statsArgsForCall = append(fake.statsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.StatsStub
	fakeReturns := fake.statsReturns
	fake.recordInvocation("Stats", []interface{}{arg1, arg2})
	fake.statsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDriver) StatsCallCount() int {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return len(fake.statsArgsForCall)
}

func (fake *FakeDriver) StatsCalls(stub func(lager.Logger, string) (groot.VolumeStats, error)) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = stub
}

func (fake *FakeDriver) StatsArgsForCall(i int) (lager.Logger, string) {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	argsForCall := fake.statsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDriver) StatsReturns(result1 groot.VolumeStats, result2 error) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	fake.statsReturns = struct {
		result1 groot.VolumeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeDriver) StatsReturnsOnCall(i int, result1 groot.VolumeStats, result2 error) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	if fake.statsReturnsOnCall == nil {
		fake.statsReturnsOnCall = make(map[int]struct {
			result1 groot.VolumeStats
			result2 error
		})
	}
	fake.statsReturnsOnCall[i] = struct {
		result1 groot.VolumeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeDriver) Unpack(arg1 lager.Logger, arg2 string, arg3 []string, arg4 io.Reader) (int64, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.unpackMutex.Lock()
	ret, specificReturn := fake.unpackReturnsOnCall[len(fake.unpackArgsForCall)]
	fake.unpackArgsForCall = append(fake.unpackArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 []string
		arg4 io.Reader
	}{arg1, arg2, arg3Copy, arg4})
	stub := fake.UnpackStub
	fakeReturns := fake.unpackReturns
	fake.recordInvocation("Unpack", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.unpackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDriver) UnpackCallCount() int {
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	return len(fake.unpackArgsForCall)
}

func (fake *FakeDriver) UnpackCalls(stub func(lager.Logger, string, []string, io.Reader) (int64, error)) {
	fake.unpackMutex.Lock()
	defer fake.unpackMutex.Unlock()
	fake.UnpackStub = stub
}

func (fake *FakeDriver) UnpackArgsForCall(i int) (lager.Logger, string, []string, io.Reader) {
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	argsForCall := fake.unpackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDriver) UnpackReturns(result1 int64, result2 error) {
	fake.unpackMutex.Lock()
	defer fake.unpackMutex.Unlock()
	fake.UnpackStub = nil
	fake.unpackReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeDriver) UnpackReturnsOnCall(i int, result1 int64, result2 error) {
	fake.unpackMutex.Lock()
	defer fake.unpackMutex.Unlock()
	fake.UnpackStub = nil
	if fake.unpackReturnsOnCall == nil {
		fake.unpackReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.unpackReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeDriver) WriteMetadata(arg1 lager.Logger, arg2 string, arg3 groot.ImageMetadata) error {
	fake.writeMetadataMutex.Lock()
	ret, specificReturn := fake.writeMetadataReturnsOnCall[len(fake.writeMetadataArgsForCall)]
	fake.writeMetadataArgsForCall = append(fake.writeMetadataArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 groot.ImageMetadata
	}{arg1, arg2, arg3})
	stub := fake.WriteMetadataStub
	fakeReturns := fake.writeMetadataReturns
	fake.recordInvocation("WriteMetadata", []interface{}{arg1, arg2, arg3})
	fake.writeMetadataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDriver) WriteMetadataCallCount() int {
	fake.writeMetadataMutex.RLock()
	defer fake.writeMetadataMutex.RUnlock()
	return len(fake.writeMetadataArgsForCall)
}

func (fake *FakeDriver) WriteMetadataCalls(stub func(lager.Logger, string, groot.ImageMetadata) error) {
	fake.writeMetadataMutex.Lock()
	defer fake.writeMetadataMutex.Unlock()
	fake.WriteMetadataStub = stub
}

func (fake *FakeDriver) WriteMetadataArgsForCall(i int) (lager.Logger, string, groot.ImageMetadata) {
	fake.writeMetadataMutex.RLock()
	defer fake.writeMetadataMutex.RUnlock()
	argsForCall := fake.writeMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDriver) WriteMetadataReturns(result1 error) {
	fake.writeMetadataMutex.Lock()
	defer fake.writeMetadataMutex.Unlock()
	fake.WriteMetadataStub = nil
	fake.writeMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) WriteMetadataReturnsOnCall(i int, result1 error) {
	fake.writeMetadataMutex.Lock()
	defer fake.writeMetadataMutex.Unlock()
	fake.WriteMetadataStub = nil
	if fake.writeMetadataReturnsOnCall == nil {
		fake.writeMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bundleMutex.RLock()
	defer fake.bundleMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	fake.writeMetadataMutex.RLock()
	defer fake.writeMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDriver) recordInvocation(key string, args []interface{}) {
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

var _ groot.Driver = new(FakeDriver)
