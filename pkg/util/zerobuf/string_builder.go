package zerobuf

// Builder
// compatible with bytes.Buffer
// compatible with strings.Builder
// compatible with "go.uber.org/zap/buffer".Buffer
// compatible with "github.com/dubbogo/gost/bytes".Buffer
type Builder interface {
	WriteString(string) (int, error)
	WriteByte(byte) error
	Write([]byte) (int, error)
	Len() int
	Cap() int
	String() string
	Reset()
}

type Stringer interface {
	Builder
	// Deprecated no-safe
	UnsafeString() string
}

type Byteser interface {
	Builder
	Bytes() []byte
}
