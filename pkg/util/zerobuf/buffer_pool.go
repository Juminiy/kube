package zerobuf

import "github.com/Juminiy/kube/pkg/util"

type Allocator interface {
	Grow(int)
}

// Pooler
// can use Free() by
// defer Pooler.Free()
type Pooler interface {
	Free()
}

type Kind int

const (
	Default Kind = iota
	Tiny
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

func Get(size ...Kind) String {
	sz := Default
	if len(size) > 0 {
		sz = size[0]
	}
	return getPool(sz)
}

func getPool(size Kind) String {
	switch size {
	case Tiny:
		return tinyPool.get()

	case Small:
		return smallPool.get()

	case Medium:
		return mediumPool.get()

	case Large:
		return largePool.get()

	case ExtraLarge:
		return extraLargePool.get()

	default:
		return tinyPool.get()
	}
}
