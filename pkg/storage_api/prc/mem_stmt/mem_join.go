package mem_stmt

import (
	"github.com/samber/lo"
	"slices"
)

type Stmt struct {
}

type Join struct {
}

func (j Join) check(r [][]string) (rC, cC int) {
	if len(r) == 0 {
		return 0, 0
	}
	return len(r), len(r[0])
}

func (j Join) valid(r [][]string) bool {
	rC, cC := j.check(r)
	return rC > 0 && cC > 0
}

func (j Join) Cross(r1, r2 [][]string) [][]string {
	if !j.valid(r1) || !j.valid(r2) {
		return nil
	}
	r12p := make([][]string, 0, len(r1)*len(r2))
	for iIdx := range r1 {
		for jIdx := range r2 {
			r12p = append(r12p, append(slices.Clone(r1[iIdx]), r2[jIdx]...))
		}
	}
	return r12p
}

func (j Join) Inner(r1, r2 [][]string, r1ci, r2ci int) [][]string {
	if !j.valid(r1) || !j.valid(r2) {
		return nil
	}
	_, r1c := j.check(r1)
	_, r2c := j.check(r2)
	if r1c <= r1ci || r2c <= r2ci {
		return nil
	}
	r12p := make([][]string, 0, min(len(r1), len(r2)))
	for iIdx := range r1 {
		for jIdx := range r2 {
			if r1[iIdx][r1ci] == r2[jIdx][r2ci] {
				r12p = append(r12p, append(slices.Clone(r1[iIdx]), r2[jIdx]...))
			}
		}
	}
	return r12p
}

func (j Join) Left(r1, r2 [][]string, r1ci, r2ci int) [][]string {
	if !j.valid(r1) || !j.valid(r2) {
		return nil
	}
	_, r1c := j.check(r1)
	_, r2c := j.check(r2)
	if r1c <= r1ci || r2c <= r2ci {
		return nil
	}
	r12p := make([][]string, 0, min(len(r1), len(r2)))
	for iIdx := range r1 {
		var iOk bool
		for jIdx := range r2 {
			ijOk := r1[iIdx][r1ci] == r2[jIdx][r2ci]
			if ijOk {
				r12p = append(r12p, append(slices.Clone(r1[iIdx]), r2[jIdx]...))
			}
			iOk = iOk || ijOk
		}
		if !iOk {
			r12p = append(r12p, append(slices.Clone(r1[iIdx]), nullStringSlice(r2c)...))
		}
	}
	return r12p
}

func (j Join) Right(r1, r2 [][]string, r1ci, r2ci int) [][]string {
	return j.Left(r2, r1, r2ci, r1ci)
}

func nullStringSlice(n int) []string {
	return lo.Times(n, func(_ int) string {
		return "NULL"
	})
}
