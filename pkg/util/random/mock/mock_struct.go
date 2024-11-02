package mock

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
)

func Struct(v any) {
	itv := indir(v)
	fieldTyp := itv.StructFieldsType()
	fieldTagKv := itv.StructParseTagKV(mockTag)

	for name, typ := range fieldTyp {
		ruledval := newRule(fieldTagKv[name]).value()
		kind := tKind(typ.Kind())

		switch {
		case isMeta(kind):
			itv.StructSetFields(map[string]any{name: ruledval[kind]})

		case typ == _timeTyp:
			itv.StructSetFields(map[string]any{name: defaultTime()})

		default:
			stdlog.WarnF("do not support type: %s", typ.String())
		}

	}

}
