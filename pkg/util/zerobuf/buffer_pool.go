package zerobuf

import "github.com/Juminiy/kube/pkg/util"

type Allocator interface {
	Grow(int)
}

type Pooler interface {
	Free()
}

const (
	Tiny = iota + 1
	Small
	Medium
	Large
	ExtraLarge

	size256B = 1 << 8       // 8 -> TinyPool
	size1K   = 1 * util.Ki  // 10 -> SmallPool
	size16K  = 16 * util.Ki // 14 -> MediumPool
	size1M   = util.Mi      // 20 -> LargePool
	size16M  = 16 * util.Mi // 24 -> ExtraLargePool
)

var (
	tinyPool       = newApBufferPool(size256B)
	smallPool      = newApBufferPool(size1K)
	mediumPool     = newApBufferPool(size16K)
	largePool      = newApBufferPool(size1M)
	extraLargePool = newApBufferPool(size16M)
)

func Get() String {
	return newApBufferPool(size256B).get()
}
