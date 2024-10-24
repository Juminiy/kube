package zero_reflect

import (
	"github.com/modern-go/reflect2"
	"testing"
)

func TestGetSet(t *testing.T) {
	var i int = 114
	var j int = 514
	reflect2.TypeOf(i).Set(&i, &j)
	t.Log(i, j)
}
