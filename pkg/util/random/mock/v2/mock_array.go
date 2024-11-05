// Package mockv2 was generated
package mockv2

import ()

func Array(v any) {
	indirTv := indir(v)
	if !indirTv.ArrayCanSet() {
		return
	}

	sliceOrArraySet(indirTv)
}
