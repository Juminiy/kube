package service

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"strings"
	"sync"
)

var _clipBoard = sync.Map{}

func ClipList() []string {
	s := make([]string, 0, util.MagicSliceCap)
	_clipBoard.Range(func(key, value any) bool {
		s = append(s, key.(string))
		return true
	})
	return s
}

func ClipAdd(s string) {
	if len(s) > 0 {
		_clipBoard.Store(s, struct{}{})
	}
}

func ClipDel(s string) {
	if len(s) > 0 {
		_clipBoard.Delete(s)
	}
}

func ClipSearch(s string) []string {
	return lo.FilterMap(ClipList(), func(item string, _ int) (string, bool) {
		if strings.Contains(item, s) || strings.Contains(s, item) {
			return item, true
		}
		return "", false
	})
}
