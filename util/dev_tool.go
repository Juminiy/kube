package util

type Page struct {
	code *int64 // equal to pageNumberOf
	size *int64 // equal to pageSizeOf
	offs *int64 // equal to recordOffsetOf
}

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
