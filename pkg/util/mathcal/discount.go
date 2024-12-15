package mathcal

import (
	"sort"

	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
)

const (
	PaymentCNY = "¥"
	PaymentUSD = "$"

	CountNumber = "个"
	CountWeight = "克"

	OrderByTotal = "total"
	OrderByCount = "count"
	OrderByUnit  = "unit"
)

type DiscountHdr struct {
	Payment    string // PaymentCNY, PaymentUSD
	TotalPrice []float64
	Count      string // CountNumber, CountWeight
	TotalCount []int
	OrderBy    string // OrderByTotal, OrderByCount, OrderByUnit

	grp    *discountGroup
	len    int
	sorted bool
}

func NewCNYNum(tot []float64, cnt []int) *DiscountHdr {
	return &DiscountHdr{
		Payment:    PaymentCNY,
		TotalPrice: tot,
		Count:      CountNumber,
		TotalCount: cnt,
		OrderBy:    OrderByUnit,
	}
}

func NewCNYWeight(tot []float64, cnt []int) *DiscountHdr {
	return &DiscountHdr{
		Payment:    PaymentCNY,
		TotalPrice: tot,
		Count:      CountWeight,
		TotalCount: cnt,
		OrderBy:    OrderByUnit,
	}
}

func (h *DiscountHdr) Cal() *DiscountHdr { h.cal(); return h }

func (h *DiscountHdr) cal() {
	h.sorted = false
	if ltot, lcnt := len(h.TotalPrice), len(h.TotalCount); ltot != lcnt ||
		lcnt == 0 {
		util.DevPanic("invalid input TotalPrice, TotalCount")
	} else {
		h.len = lcnt
	}

	if !util.ElemIn(h.Payment, PaymentCNY, PaymentUSD) {
		h.Payment = PaymentCNY
	}
	if !util.ElemIn(h.Count, CountWeight, CountNumber) {
		util.DevPanic("invalid input Count Method, unit number(s) or weight(g)")
	}
	if !util.ElemIn(h.OrderBy, OrderByTotal, OrderByUnit, OrderByCount) {
		h.OrderBy = OrderByUnit
	}

	h.grp = &discountGroup{
		hdr:       h,
		discounts: make([]discount, h.len),
	}
	for i := range h.len {
		h.grp.setIndex(i, h.TotalPrice[i], h.TotalCount[i])
	}
}

func (h *DiscountHdr) Sort() *DiscountHdr {
	h.sorted = true
	sort.Sort(h.grp)
	h.grp.setRate()
	return h
}

func (h *DiscountHdr) Log() {
	for i := range h.len {
		stdlog.InfoF("%d, %s%.2f/%d%s = %s%.4f, 500g = %s%.2f, discount_0 = %.2f%%, discount_1 = %.2f%%",
			h.grp.discounts[i].index+1,
			h.Payment, h.grp.discounts[i].total, h.grp.discounts[i].count, h.Count, // original data
			h.Payment, h.grp.discounts[i].unit, // unit price
			h.Payment, h.grp.discounts[i].g500Price, // 500g price
			h.grp.discounts[i].rateCmpNext, h.grp.discounts[i].rateCmpExpensive, // discount rate
		)
	}
}

type discount struct {
	index int
	total float64
	count int

	unit      float64
	g500Price float64

	rateCmpNext      float64
	rateCmpExpensive float64
}

type discountGroup struct {
	hdr       *DiscountHdr
	discounts []discount
}

func (s discountGroup) setIndex(i int, total float64, count int) {
	unit := total / float64(count)
	s.discounts[i].index = i
	s.discounts[i].total = total
	s.discounts[i].count = count
	s.discounts[i].unit = unit
	if s.hdr.Count == CountWeight {
		s.discounts[i].g500Price = unit * 500
	}
}

func (s discountGroup) setRate() {
	if !s.hdr.sorted {
		return
	}

	var mostExpensive = s.discounts[s.Len()-1]
	for i := 0; i < s.Len()-1; i++ {
		cur, next := s.discounts[i], s.discounts[i+1]
		s.discounts[i].rateCmpNext = ((next.unit / cur.unit) - 1) * 100
		s.discounts[i].rateCmpExpensive = ((mostExpensive.unit / cur.unit) - 1) * 100
	}
}

func (s discountGroup) Swap(i, j int) {
	s.discounts[i], s.discounts[j] = s.discounts[j], s.discounts[i]
}

func (s discountGroup) Less(i, j int) bool {
	switch s.hdr.OrderBy {
	case OrderByTotal:
		return s.discounts[i].total < s.discounts[j].total

	case OrderByCount:
		return s.discounts[i].count < s.discounts[j].count

	case OrderByUnit:
		return s.discounts[i].unit < s.discounts[j].unit

	default:
		stdlog.Error("unsupported order by")
		return false
	}
}

func (s discountGroup) Len() int {
	return s.hdr.len
}
