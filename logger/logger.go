package logger

import (
	"fmt"
	"io"
	"log"
	"time"
)

// LogMode specifies pre-defined log entry format
type LogMode int

/* About Mode
 * 2 modes are available for now
 * [1] Datadog compatible mode
 *     "[DEBUG]	2020-03-25T07:49:06.648Z	[Test] test write to file"
 * [2] Standard simple mode
 *     "2020/03/25 16:52:16 [DEBUG][Test] test write to file"
 */

const (
	DdgMode LogMode = 1 << iota
	StdMode
)

const (
	// Format: [Level], [Prefix], [Title], message
	ddgFormat = "[%s]\t%s\t[%s] %s\n"
	stdFormat = "[%s]%s[%s] %s\n"
)

const (
	labelInfo  = "INFO"
	labelDebug = "DEBUG"
	labelWarn  = "WARN"
	labelFatal = "FATAL"
	labelError = "ERROR"
)

type Logger struct {
	IsDebug bool
	mode    LogMode
	format  string

	lwarn  string
	lerror string
	linfo  string
	lfatal string
	ldebug string

	ilogger *log.Logger
}

func (logger *Logger) SetLogger(ilogger *log.Logger) {
	logger.ilogger = ilogger
}

func (logger *Logger) SetOutput(out io.Writer) {
	logger.ilogger.SetOutput(out)
}

func (logger *Logger) SetFlags(flag int) {
	logger.ilogger.SetFlags(flag)
}

func (logger *Logger) SetLabels(linfo, ldebug, lwarn, lfatal, lerror string) {
	logger.linfo = linfo
	logger.ldebug = ldebug
	logger.lwarn = lwarn
	logger.lfatal = lfatal
	logger.lerror = lerror
}

func (logger *Logger) SetDdgMode() {
	logger.mode = DdgMode
	logger.format = ddgFormat
	logger.ilogger.SetFlags(0)
}

func (logger *Logger) SetStdMode() {
	logger.mode = StdMode
	logger.format = stdFormat
	logger.ilogger.SetFlags(log.LstdFlags)
}

func (logger *Logger) GetMode() LogMode {
	return logger.mode
}

func (logger *Logger) Print(v ...interface{}) {
	logger.ilogger.Print(v...)
}

func (logger *Logger) Printf(format string, v ...interface{}) {
	logger.ilogger.Printf(format, v...)
}

func (logger *Logger) Println(v ...interface{}) {
	logger.ilogger.Println(v...)
}

func (logger *Logger) Output(i int, s string) error {
	return logger.ilogger.Output(i, s)
}

func (logger *Logger) prefix() string {
	if logger.mode == DdgMode {
		return time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	}
	return ""
}

func (logger *Logger) Warning(title, format string, v ...interface{}) {
	logger.Printf(logger.format, logger.lwarn, logger.prefix(), title, fmt.Sprintf(format, v...))
}

func (logger *Logger) Info(title, format string, v ...interface{}) {
	logger.Printf(logger.format, logger.linfo, logger.prefix(), title, fmt.Sprintf(format, v...))
}

func (logger *Logger) Fatal(title, format string, v ...interface{}) {
	logger.Printf(logger.format, logger.lfatal, logger.prefix(), title, fmt.Sprintf(format, v...))
}

func (logger *Logger) Error(title, format string, v ...interface{}) {
	logger.Printf(logger.format, logger.lerror, logger.prefix(), title, fmt.Sprintf(format, v...))
}

func (logger *Logger) Debug(title, format string, v ...interface{}) {
	if !logger.IsDebug {
		return
	}
	logger.Printf(logger.format, logger.ldebug, logger.prefix(), title, fmt.Sprintf(format, v...))
}

func NewLogger(debug bool, out io.Writer) *Logger {
	return &Logger{
		IsDebug: debug,
		ilogger: log.New(out, "", log.LstdFlags),
		mode:    StdMode,
		format:  stdFormat,
		linfo:   labelInfo,
		ldebug:  labelDebug,
		lwarn:   labelWarn,
		lfatal:  labelFatal,
		lerror:  labelError,
	}
}

func NewDdgLogger(debug bool, out io.Writer) *Logger {
	return &Logger{
		IsDebug: debug,
		ilogger: log.New(out, "", 0),
		mode:    DdgMode,
		format:  ddgFormat,
		linfo:   labelInfo,
		ldebug:  labelDebug,
		lwarn:   labelWarn,
		lfatal:  labelFatal,
		lerror:  labelError,
	}
}
