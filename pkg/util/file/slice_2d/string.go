package slice_2d

import (
	"github.com/Juminiy/kube/pkg/util"
	"golang.org/x/net/context"
	"k8s.io/apimachinery/pkg/util/sets"
)

type String[Model any] struct {
	NameRequired   sets.Set[string]
	HeaderToName   map[string]string
	DefaultConvert map[string]Value
	StaticConvert  map[string]Value
	DynamicConvert map[string]Value
	FieldUnique    map[string]sets.Set[string]
	List           *[]Model
	Actions        []func() error

	Headers  []string
	Cells    []string
	warnings *util.ErrHandle
	errors   []error
}

func (s String[Model]) Valid() bool {
	return true
}

func (s String[Model]) Do(ctx context.Context) {

}

type Value struct {
	Convert func(string) string
	Valid   func(string) bool
}
