package mockv2

import (
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

func TestSlice(t *testing.T) {
	t0sl := make([]t0, 32)
	Slice(t0sl)
	t.Log(len(safe_json.String(t0sl)))
}
