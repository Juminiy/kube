package mathcal

import "testing"

func TestDisCount0(t *testing.T) {
	tot := []float64{3.8, 6.2, 9.9, 13.8, 16.9, 19.9, 22.9, 26.5, 29.9, 33.5, 36.9}
	cnt := []int{4, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	NewCNYNum(tot, cnt).Cal().Sort().Log()
}

func TestDisCount1(t *testing.T) {
	tot := []float64{5.88, 8.88, 14.88, 19.88, 8.98, 15.88, 20.88, 9.98, 19.88, 26.88, 13.88, 15.88, 17.88, 17.88, 25.88}
	cnt := []int{115, 230, 230 * 2, 230 * 3, 230, 230 * 2, 230 * 3, 310, 310 * 2, 310 * 3, 115 + 230, 230 + 230, 310 + 230, 230 + 310, 310*2 + 230}
	NewCNYWeight(tot, cnt).Cal().Sort().Log()
}

func discountPerLog(t *testing.T, tot []float64, cnt []int, per string) {
	for i := range tot {
		t.Logf("¥%.2f/%d%s = %.4f, 500g = ¥%.2f", tot[i], cnt[i], per, tot[i]/float64(cnt[i]), tot[i]/float64(cnt[i])*500)
	}
}

func TestLogLog(t *testing.T) {
	t.Log(LoopLog(1e10, 2))
	t.Log(LogLog(1e10, 2, 2))
}
