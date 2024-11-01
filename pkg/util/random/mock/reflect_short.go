package mock

import "github.com/Juminiy/kube/pkg/util/safe_reflect"

var indir = safe_reflect.IndirectOf

const (
	mockTag = "mock"
)
