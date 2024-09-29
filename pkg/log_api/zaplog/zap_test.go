//go:build unix

package zaplog

import (
	"path/filepath"
	"testing"
)

var _cfg = New().
	WithLogEngineSugar().
	WithLogLevel("info").
	WithLogCaller(false).
	WithLogStackTrace(false).
	WithOutputPaths(filepath.Join(testDir, "app.log")).
	WithErrorOutputPaths(filepath.Join(testDir, "app_error.log"))

func TestZapLog(t *testing.T) {
	_cfg.Load()
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

/*// Once Tested, Comment them immediately with (Shift+Ctrl+/)
func TestStdLogFatal(t *testing.T) {
	cfg.Load()
	Fatal("1", "2", "3")
}

func TestStdLogFatalF(t *testing.T) {
	cfg.Load()
	FatalF("ex %d", 1)
}

func TestStdLogFatalW(t *testing.T) {
	cfg.Load()
	FatalW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}

func TestStdLogPanic(t *testing.T) {
	cfg.Load()
	Panic("1", "2", "3")
}

func TestStdLogPanicF(t *testing.T) {
	cfg.Load()
	PanicF("ex %d", 1)
}

func TestStdLogPanicW(t *testing.T) {
	cfg.Load()
	PanicW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}
*/
