package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util"
	"maps"
	"sync"
)

type Config struct {
	// Tag
	// tag scope to valid struct field
	// `tag:"k1:v1;k2:v2;k3:v3"`
	Tag string

	// OnErrorStop
	// stop check immediately when tag error or valid error
	// return false when call Any, Array, Map, Slice, Struct
	// return error when call Any, ArrayE, MapE, SliceE, StructE
	OnErrorStop bool

	// IgnoreTagError
	// skip check field when tag format error, raise no error
	IgnoreTagError bool

	// IndirectValue
	// can apply field tag validator to pointer value
	IndirectValue bool

	// AllowEmbedStruct
	AllowEmbedStruct bool

	// FloatPrecision
	// set default precision by -1
	// for example: set 6 is after point six bit
	FloatPrecision int

	once  sync.Once
	apply tagApplyKindT
}

var _defaultConfig = &Config{
	Tag:              "valid",
	OnErrorStop:      false,
	IgnoreTagError:   false,
	IndirectValue:    false,
	AllowEmbedStruct: false,
	FloatPrecision:   -1,
	once:             sync.Once{},
}

func Default() *Config {
	return _defaultConfig.Load()
}

var _strictConfig = &Config{
	Tag:              "valid",
	OnErrorStop:      false,
	IgnoreTagError:   false,
	IndirectValue:    true,
	AllowEmbedStruct: true,
	FloatPrecision:   66,
	once:             sync.Once{},
}

func Strict() *Config { return _strictConfig.Load() }

func (cfg *Config) Load() *Config {
	cfg.once.Do(func() {
		cfg.apply = maps.Clone(_apply)
		if cfg.IndirectValue {
			for tag, kinds := range cfg.apply {
				cfg.apply[tag] = util.MapInsert(maps.Clone(kinds), kPtr)
			}
		}
	})
	return cfg
}
