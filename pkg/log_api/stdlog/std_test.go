package stdlog

import (
	kubeinternal "github.com/Juminiy/kube/pkg/internal_api"
	"path/filepath"
	"testing"
)

var testDir, testDirErr = kubeinternal.GetWorkPath("testdata", "test_log")

var stdCfg = New().
	WithLogPath(filepath.Join(testDir, "console.log"))

func testInit() {
	stdCfg.Load()
	if testDirErr != nil {
		panic(testDirErr)
	}
}

// +passed +windows +darwin
func TestStdLog(t *testing.T) {
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

/*// +passed +windows +darwin
func TestStdLogFatal(t *testing.T) {
	testInit()
	Fatal("1", "2", "3")
}

// +passed +windows +darwin
func TestStdLogFatalF(t *testing.T) {
	testInit()
	FatalF("ex %d", 1)
}

// +passed +windows +darwin
func TestStdLogFatalW(t *testing.T) {
	testInit()
	FatalW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}

// +passed +windows +darwin
func TestStdLogPanic(t *testing.T) {
	testInit()
	Panic("1", "2", "3")
}

// +passed +windows +darwin
func TestStdLogPanicF(t *testing.T) {
	testInit()
	PanicF("ex %d", 1)
}

// +passed +windows +darwin
func TestStdLogPanicW(t *testing.T) {
	testInit()
	PanicW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}
*/