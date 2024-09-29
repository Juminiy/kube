package stdlog

import (
	"path/filepath"
	"testing"
)

var stdCfg = New().
	WithLogPath(filepath.Join(testDir, "console.log"))

func TestStdLog(t *testing.T) {
	stdCfg.Load()
	Debug("1", "2", "3")
	DebugF("ex %d", 1)
	DebugW("xe", "k", "v", "k2", 2, "k3", map[string]string{})

	Info("1", "2", "3")
	InfoF("ex %d", 1)
	InfoW("xe", "k", "v", "k2", 2, "k3", map[string]string{})

	Warn("1", "2", "3")
	WarnF("ex %d", 1)
	WarnW("xe", "k", "v", "k2", 2, "k3", map[string]string{})

	Error("1", "2", "3")
	ErrorF("ex %d", 1)
	ErrorW("xe", "k", "v", "k2", 2, "k3", map[string]string{})

}

// Once Tested, Comment them immediately with (Shift+Ctrl+/)
func TestStdLogFatal(t *testing.T) {
	stdCfg.Load()
	Fatal("1", "2", "3")
}

func TestStdLogFatalF(t *testing.T) {
	stdCfg.Load()
	FatalF("ex %d", 1)
}

func TestStdLogFatalW(t *testing.T) {
	stdCfg.Load()
	FatalW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}

func TestStdLogPanic(t *testing.T) {
	stdCfg.Load()
	Panic("1", "2", "3")
}

func TestStdLogPanicF(t *testing.T) {
	stdCfg.Load()
	PanicF("ex %d", 1)
}

func TestStdLogPanicW(t *testing.T) {
	stdCfg.Load()
	PanicW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}
