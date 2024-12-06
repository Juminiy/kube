package safe_validator

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"math/rand/v2"
	"testing"
)

var correctT0Elem = [3]t0{
	{
		I0:   3,
		F0:   0.01,
		S0:   "12345",
		IPtr: util.New(2),
		SPtr: util.New("c"),
		Arr0: []int{1},
		Map0: map[string]string{"k1": "v1"},
		I1:   22,
	},
	{
		I0:   1,
		F0:   -0.01,
		S0:   "7834122",
		IPtr: util.New(1),
		SPtr: util.New("a"),
		Arr0: []int{2},
		Map0: map[string]string{"srv6": random.AllString(10)},
		I1:   99,
	},
	{
		I0:   2,
		F0:   0.09,
		S0:   random.NumericString(5),
		IPtr: util.New(2),
		SPtr: util.New("b"),
		Arr0: []int{88, 99211},
		Map0: map[string]string{"k9": "v1111"},
		I1:   66,
	},
}

var errT0Elem = [3]t0{
	{
		I0: rand.Int(),
		F0: rand.Float64(),
		S0: random.AllString(20),
	},
	{
		I0:   1,
		F0:   -0.01,
		S0:   "7834122",
		IPtr: util.New(1),
		SPtr: util.New("a"),
		Arr0: []int{2, 5, 6, 7},
		Map0: map[string]string{
			"k1": "",
			"k2": "",
			"k3": "",
			"k4": "",
			"k5": "",
			"k6": "",
			"k7": "",
			"k8": "",
		},
		I1: 99,
	},
	{},
}

func TestStrict_Slice(t *testing.T) {
	t.Log(Strict().SliceE(correctT0Elem[:]))
	t.Log(Strict().SliceE(errT0Elem[:]))
}

func TestStrict_StructBug(t *testing.T) {
	t.Log(Strict().StructE(struct {
		F0 float64 `valid:"not_zero;range:-0.1~0.1;enum:-0.01,0.01,0.09"`
		F1 float64 `valid:"not_zero;range:-0.1~0.1;enum:-0.01,0.01,0.09"`
	}{
		F0: 666,
		F1: 0.01,
	}))
}
