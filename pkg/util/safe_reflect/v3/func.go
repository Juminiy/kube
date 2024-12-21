package safe_reflectv3

func (v V) CallMethod(name string, args []any) (rets []any, called bool) {
	method := v.MethodByName(name)
	if method == ZeroValue() {
		return
	}
	return Wrap(method).FuncCall(args)
}

func (tv Tv) FuncCall(in []any) (out []any, called bool) {
	numIn := tv.NumIn()
	if numIn != len(in) {
		return
	}
	rtin := rts(in)
	for i := range numIn {
		if rtin[i] != tv.In(i) {
			return
		}
	}
	return Anys(tv.Call(rvs(in))), true
}
