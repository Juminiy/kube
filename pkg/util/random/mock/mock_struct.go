package mock

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"reflect"
)

func Struct(v any) {
	cacheKey, tv, ok := structCache(v)
	structSet(tv.val, tv.typ)
	if !ok {
		cachePut(cacheKey, tv.typ)
	}
}

func structCache(v any) (
	uintptr, // cacheKey
	tStructTv, // cachedValue
	bool, // ifCached
) {
	indirTv := indir(v)
	cacheKey, cacheVal := cacheGet(v)
	if cacheVal != nil {
		return cacheKey, tStructTv{
			val: indirTv,
			typ: cacheVal.(*tStructTyp),
		}, true
	}
	return cacheKey, tStructTv{
		val: indirTv,
		typ: &tStructTyp{
			FieldTyp:   indirTv.StructFieldsType(),
			FieldTagKv: indirTv.StructParseTagKV(mockTag),
			FieldRule:  make(map[string]*Rule, indirTv.FieldLen()),
		},
	}, false
}

func structCached(t reflect.Type) *tStructTyp {
	_, cacheVal := cacheByTyp(t)
	if cacheVal != nil {
		return cacheVal.(*tStructTyp)
	}
	return nil
}

func structSet(indirTv safe_reflect.TypVal, structOf *tStructTyp) {
	for name, typ := range structOf.FieldTyp {
		tagkv := structOf.FieldTagKv[name]
		if util.MapOk(tagkv, "null") { // skip null
			continue
		}
		kind := tKind(typ.Kind())
		var fieldValue any

		// for special
		fieldValue = specialRule(tagkv)

		// for rule
		if fieldValue == nil {
			fieldrule, ok := structOf.FieldRule[name]
			if !ok {
				fieldrule = newRule(tagkv).parse()
				structOf.FieldRule[name] = fieldrule
			}
			fieldValue = fieldrule.value()[kind]
		}

		switch {
		case kind.isMeta():
			indirTv.StructSetFields(map[string]any{name: fieldValue})

		case typ == _timeTyp:
			indirTv.StructSetFields(map[string]any{name: defaultTime()})

		default:
			stdlog.WarnF("do not support type: %s", typ.String())
		}

	}
}

type tStructTv struct {
	// by indir
	val safe_reflect.TypVal // no cached
	// struct typ info
	typ *tStructTyp // cached
}

type tStructTyp struct {
	FieldTyp   map[string]reflect.Type
	FieldTagKv safe_reflect.FieldTagKV
	FieldRule  map[string]*Rule
}
