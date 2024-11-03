package zero_reflect

import (
	"github.com/Juminiy/kube/pkg/internal"
	"sync"
)

type ConfigOption struct {
	_      internal.NoCmp
	noCopy internal.NoCopy
	sync.Once
}

func New() *ConfigOption {
	return &ConfigOption{}
}

func (o *ConfigOption) Load() {
	o.Do(Init)
}

func (o *ConfigOption) WithNoPointer() *ConfigOption {
	_noPointerLevel = _noPointer
	return o
}

func (o *ConfigOption) WithStructSpec() *ConfigOption {
	_noPointerLevel = _structSpec
	return o
}

func (o *ConfigOption) WithComparable() *ConfigOption {
	_noPointerLevel = _mustComparable
	return o
}
