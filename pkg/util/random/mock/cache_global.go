package mock

import (
	runtimeutil "github.com/Juminiy/kube/pkg/util/runtime"
	"sync"
)

var _global = sync.Map{}

func cacheGet(v any) (uintptr, any) {
	vTypKey := runtimeutil.EFaceOf(v).Type()
	vInfoCache, ok := _global.Load(vTypKey)
	if ok {
		return vTypKey, vInfoCache
	}
	return vTypKey, nil
}

func cachePut(k uintptr, v any) {
	_global.Store(k, v)
}

func cacheByTyp(t any) (uintptr, any) {
	tValKey := runtimeutil.EFaceOf(t).Value()
	vInfoCache, ok := _global.Load(tValKey)
	if ok {
		return tValKey, vInfoCache
	}
	return tValKey, nil
}
