package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestHasField(t *testing.T) {
	t.Log(HasField(t0{}, "F0", "")) // has name and type is
	t.Log(HasField(t0{}, "F1", "")) // has name but type not
	t.Log(HasField(t0{}, "F2", "")) // no name
	t.Log(HasField(t0{}, "F3", 1))  // others
}

func TestSetField(t *testing.T) {
	tval := t0{}
	t.Log(tval)

	SetField(tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(tval, "F3", "set F3") // others
	t.Log(tval)

	SetField(&tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(&tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(&tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(&tval, "F3", "set F3") // others
	t.Log(tval)
}

func TestHasFields(t *testing.T) {
	t.Log(HasFields(t0{}, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": 1,                     // has name but type not
	}))
}

func TestSetFields(t *testing.T) {
	tval := t0{}

	SetFields(tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)
}

func TestHasField2(t *testing.T) {
	sl := []t0{{}, {}, {}}
	t.Log(HasField(sl, "F0", "")) // has name and type is
	t.Log(HasField(sl, "F1", "")) // has name but type not
	t.Log(HasField(sl, "F2", "")) // no name
	t.Log(HasField(sl, "F3", 1))  // others

	// nil
	sl = nil
	t.Log(HasField(sl, "F0", "")) // has name and type is
	t.Log(HasField(sl, "F1", "")) // has name but type not
	t.Log(HasField(sl, "F2", "")) // no name
	t.Log(HasField(sl, "F3", 1))  // others
}

func TestSetField2(t *testing.T) {
	tval := []t0{{}, {}, {}}
	t.Log(tval)

	SetField(tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(tval, "F3", "set F3") // others
	t.Log(tval)

	tval = nil
	SetField(&tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(&tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(&tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(&tval, "F3", "set F3") // others
	t.Log(tval)

	tval = []t0{{}, {}, {}}
	SetField(&tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(&tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(&tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(&tval, "F3", "set F3") // others
	t.Log(tval)
}

func TestHasFields2(t *testing.T) {
	t.Log(HasFields([]t0{}, map[string]any{
		"F0": "ccc", // has name and type is
	}))
}

func TestSetFields2(t *testing.T) {
	tval := []t0{{}}

	SetFields(tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	tval = []t0{{}}
	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	tval = nil
	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	tval = []t0{}
	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)
}
