package util

type Page struct {
	code *int64 // equal to pageNumberOf, pageCodeOf
	size *int64 // equal to pageSizeOf, pageCapOf
	offs *int64 // equal to recordOffsetOf
}

// DefaultPage
// immutable default page config
var (
	DefaultPage = NewPageConfig(1, 10)
)

func NewPageConfig(pageNum, pageSize int64) *Page {
	return &Page{
		NewInt64(pageNum),
		NewInt64(pageSize),
		NewInt64(0),
	}
}

func (p *Page) WithPageNum(pageNum int64) *Page {
	p.code = NewInt64(pageNum)
	return p
}

func (p *Page) WithPageSize(pageSize int64) *Page {
	p.size = NewInt64(pageSize)
	return p
}

func (p *Page) WithOffsetNum(offsetNum int64) *Page {
	p.offs = NewInt64(offsetNum)
	return p
}

func (p *Page) Pair() (*int64, *int64) {
	return p.code, p.size
}

func (p *Page) Page() *int64 {
	return p.code
}

func (p *Page) Size() *int64 {
	return p.size
}

func (p *Page) PairValue() (int64, int64) {
	return *p.code, *p.size
}

func (p *Page) PageValue() int64 {
	return *p.code
}

func (p *Page) SizeValue() int64 {
	return *p.size
}

func (p *Page) PageIntValue() int {
	return int(p.PageValue())
}

func (p *Page) SizeIntValue() int {
	return int(p.SizeValue())
}
