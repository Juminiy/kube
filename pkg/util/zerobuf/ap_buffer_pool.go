package zerobuf

type apBufferPool struct {
	p *Pool[*apString]
}

func newApBufferPool(size int) apBufferPool {
	return apBufferPool{
		p: New(func() *apString {
			return &apString{
				b: make([]byte, 0, size),
			}
		}),
	}
}

func (p apBufferPool) get() *apString {
	ap := p.p.Get()
	ap.Reset()
	ap.apBufferPool = p
	return ap
}

func (p apBufferPool) put(ap *apString) {
	p.p.Put(ap)
}
