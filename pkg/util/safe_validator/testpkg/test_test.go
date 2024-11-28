package testpkg

import (
	"github.com/Juminiy/kube/pkg/util/safe_validator"
	"testing"
)

func TestTest(t *testing.T) {
	cfg := safe_validator.Config{
		Tag:            "valid_tag",
		OnErrorStop:    true,
		IgnoreTagError: false,
		IndirectValue:  true,
		FloatPrecision: 0,
	}
	t.Log(cfg.StructE(1))
}
