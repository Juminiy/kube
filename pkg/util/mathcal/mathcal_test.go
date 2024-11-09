package mathcal

import "testing"

func TestCount(t *testing.T) {
	tot := []float64{3.8, 6.2, 9.9, 13.8, 16.9, 19.9, 22.9, 26.5, 29.9, 33.5, 36.9}
	cnt := []int{4, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i := range tot {
		t.Logf("%f/%d = %f", tot[i], cnt[i], tot[i]/float64(cnt[i]))
	}
}
