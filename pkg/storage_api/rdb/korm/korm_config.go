package korm

type Config struct {
	Tag        *Tag
	QRetry     int
	UseNumber  bool
	UsePointer bool
	NoTag      *NoTag
	ConnC      *ConnC
}

type Tag struct {
	Key        string
	Column     string
	Type       string
	Size       string
	PrimaryKey string
	Unique     string
	Default    string
	Precision  string
	Scale      string
	NotNULL    string
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
	Tag: &Tag{
		Key:        "k",
		Column:     "col",
		Type:       "typ",
		Size:       "sz",
		PrimaryKey: "pk",
		Unique:     "u",
		Default:    "d",
		Precision:  "pr",
		Scale:      "sc",
		NotNULL:    "nn",
	},
	QRetry:     _DefaultQRetry,
	UseNumber:  true,
	UsePointer: true,
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
