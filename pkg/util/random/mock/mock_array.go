package mock

func Array(v any) {
	indirTv := indir(v)
	if !indirTv.ArrayCanSet() {
		return
	}

	sliceOrArraySet(indirTv)
}
