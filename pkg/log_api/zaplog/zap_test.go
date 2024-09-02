package zaplog

import "testing"

var testGlobalConfig = NewConfig().
	WithLogEngineSugar().
	WithLogLevel("info").
	WithLogCaller().
	WithLogStackTrace().
	WithOutputPaths("/home/wz/test_dir/app.log").
	WithErrorOutputPaths("/home/wz/test_dir/app_error.log")

// TODO: something seems wrong
func TestZapLog(t *testing.T) {
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
/*func TestStdLogFatal(t *testing.T) {
	Fatal("1", "2", "3")
}

func TestStdLogFatalF(t *testing.T) {
	FatalF("ex %d", 1)
}

func TestStdLogFatalW(t *testing.T) {
	FatalW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}

func TestStdLogPanic(t *testing.T) {
	Panic("1", "2", "3")
}

func TestStdLogPanicF(t *testing.T) {
	PanicF("ex %d", 1)
}

func TestStdLogPanicW(t *testing.T) {
	PanicW("xe", "k", "v", "k2", 2, "k3", map[string]string{})
}*/
