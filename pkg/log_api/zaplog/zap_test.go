package zaplog

import (
	kubeinternal "github.com/Juminiy/kube/pkg/internal"
	"path/filepath"
	"testing"
)

var testDir, testDirErr = kubeinternal.GetWorkPath("testdata", "test_log")

var _cfg = New().
	WithLogEngineSugar().
	WithLogLevel("info").
	WithLogCaller(false).
	WithLogStackTrace(false).
	WithOutputPaths(filepath.Join(testDir, "app.log")).
	WithErrorOutputPaths(filepath.Join(testDir, "app_error.log"))

func testInit() {
	_cfg.Load()
	if testDirErr != nil {
		panic(testDirErr)
	}
}

// +passed +windows
func TestZapLog(t *testing.T) {
	testInit()
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

/*// +passed +windows
func TestStdLogFatal(t *testing.T) {
	testInit()
	Fatal("1", "2", "3")
}

// +passed +windows
func TestStdLogFatalF(t *testing.T) {
	testInit()
	FatalF("ex %d", 1)
}

// +passed +windows
func TestStdLogFatalW(t *testing.T) {
	testInit()
	FatalW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}

// +passed +windows
func TestStdLogPanic(t *testing.T) {
	testInit()
	Panic("1", "2", "3")
}

// +passed +windows
func TestStdLogPanicF(t *testing.T) {
	testInit()
	PanicF("ex %d", 1)
}

// +passed +windows
func TestStdLogPanicW(t *testing.T) {
	testInit()
	PanicW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}*/
