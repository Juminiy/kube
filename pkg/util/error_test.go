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

func TestErrHandle_HasStr(t *testing.T) {
	eh := NewErrHandle()
	wg := sync.WaitGroup{}
	wg.Add(15)
	for range 5 {
		go func() { defer wg.Done(); eh.Has(nil, nil, nil) }()
		go func() { defer wg.Done(); eh.Has(ErrFaked, nil, ErrFaked) }()
		go func() { defer wg.Done(); eh.HasStr(ErrFaked.Error()) }()
	}
	wg.Wait()
	t.Log(eh.First())
	t.Log(eh.All("\t"))
	t.Log(eh.AllStr("\t"))
}

func TestErrHandle_AllStr(t *testing.T) {
	eh := NewErrHandle()
	wg := sync.WaitGroup{}
	wg.Add(9)
	for range 3 {
		go func() { defer wg.Done(); eh.Has(nil, nil, nil) }()
		go func() { defer wg.Done(); eh.Has(nil, nil, nil) }()
		go func() { defer wg.Done(); eh.HasStr("", "", "") }()
	}
	wg.Wait()
	t.Log(eh.First())
	t.Log(eh.All("\t"))
	t.Log(eh.AllStr("\t"))
}
