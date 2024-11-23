package util

import (
	"sync"
	"testing"
)

func TestNewErrHandle(t *testing.T) {
	eh := NewErrHandle()
	wg := sync.WaitGroup{}
	wg.Add(40)
	for range 20 {
		go func() { t.Log(eh.Has(nil, nil, nil)); wg.Done() }()
		go func() { t.Log(eh.Has(nil, ErrFaked, nil)); wg.Done() }()
	}
	wg.Wait()
	t.Log(eh.First())
	t.Log(eh.All())
}
