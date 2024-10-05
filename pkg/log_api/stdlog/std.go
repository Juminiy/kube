// global var
package stdlog

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/internal_api"
	"log"
	"os"
)

// same interface with zap sugar log

// xxxX
// each log a line

// xxxF
// policy: the format placeholder must:
// - use %#v in a struct type or pointer type
// - not use %v in a struct type or pointer type

// xxxW
// suggestion: W stand for with key value

const (
	logLevelDebug = "DEBUG "
	logLevelInfo  = "INFO  "
	logLevelWarn  = "WARN  "
	logLevelError = "ERROR "
	logLevelFatal = "FATAL "
	logLevelPanic = "PANIC "
)

// global config
var (
	_logPath             string
	_logTimeMicroSeconds bool
	_logCallerLongFile   bool
	_logCallerShortFile  bool
	_logTimeUTC          bool
)

// global var
var (
	_logDefaultOut = os.Stderr
	// always do not allow _logger to be nil
	// use log.Default
	// var std = New(os.Stderr, "", LstdFlags)
	_logger *log.Logger = log.New(_logDefaultOut, "", log.LstdFlags)
)

func Init() {
	lFlag := log.LstdFlags
	if _logTimeMicroSeconds {
		lFlag |= log.Lmicroseconds
	}
	if _logCallerLongFile {
		lFlag |= log.Llongfile
	} else if _logCallerShortFile {
		lFlag |= log.Lshortfile
	}
	if _logTimeUTC {
		lFlag |= log.LUTC
	}

	// use util.OSOpenFileWithCreate
	// import cycle, use func internal directly
	logFilePtr, err := internal_api.AppendCreateFile(_logPath)
	if err != nil {
		fmt.Printf("open or create stdlog filepath: %s, error: %s\n", _logPath, err.Error())
		// if any error use default by config
		_logger = log.New(_logDefaultOut, "", lFlag)
		return
	}

	_logger = log.New(logFilePtr, "", lFlag)
}
