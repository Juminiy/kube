package korm

type Config struct {
	TagKey     string
	QRetry     int
	UseNumber  bool
	UsePointer bool
	NoTag      *NoTag
	ConnC      *ConnC
}

type NoTag struct {
	ColumnSnakeCase bool
	ColumnCamelCase bool
}

type ConnC struct {
	IdleSec   int
	LifeSec   int
	IdleConns int
	OpenConns int
}

const _DefaultTagKey = `k`
const _DefaultQRetry = 3 // retry only for Query

var _DefaultConfig = &Config{
	QRetry: _DefaultQRetry,
	TagKey: _DefaultTagKey,
	NoTag: &NoTag{
		ColumnSnakeCase: true,
		ColumnCamelCase: false,
	},
	ConnC: &ConnC{
		IdleSec:   600,
		LifeSec:   600,
		IdleConns: 8,
		OpenConns: 4,
	},
}
