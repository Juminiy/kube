package stdlog

import (
	"log"
	"strings"
)

// policy: the format placeholder must:
// - use %#v in a struct type or pointer type
// - not use %v in a struct type or pointer type

// suggestion: W stand for with key value
var (
	// Debug color: purple
	Debug  = outputLine
	DebugF = outputFormatLine
	DebugW = outputKeyValueLine

	// Info color: green
	Info  = outputLine
	InfoF = outputFormatLine
	InfoW = outputKeyValueLine

	// Warn color: yellow
	Warn  = outputLine
	WarnF = outputFormatLine
	WarnW = outputKeyValueLine

	// Error color: red
	Error  = outputLine
	ErrorF = outputFormatLine
	ErrorW = outputKeyValueLine

	// Fatal color: red
	Fatal  = log.Fatalln
	FatalF = func(format string, v ...any) {
		log.Fatalf(formatLine(format), v...)
	}
	FatalW = func(msg string, kv ...any) {
		log.Print(msg)
		log.Fatalln(kv...)
	}

	// Panic color: red
	Panic  = log.Panicln
	PanicF = func(format string, v ...any) {
		log.Panicf(formatLine(format), v...)
	}
	PanicW = func(msg string, kv ...any) {
		log.Print(msg)
		log.Panicln(kv...)
	}
)

var outputLine = log.Println

func outputFormatLine(format string, v ...any) {
	log.Printf(formatLine(format), v...)
}

func outputKeyValueLine(message string, kv ...any) {
	log.Print(message)
	log.Println(kv...)
}

func formatLine(format string) string {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	return format
}
