package zero_reflect

import (
	reflect3 "github.com/goccy/go-reflect"
	"testing"
)

func TestGoccyReflect(t *testing.T) {
	reflect3.TypeOf(10)
}
