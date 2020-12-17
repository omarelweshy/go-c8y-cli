package logger

import (
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02T15:04:05.000000 MST} %{level:.5s} %{color:reset} %{message}`,
)

type NopBackend struct {
}

func (b *NopBackend) Log(level logging.Level, calldepth int, rec *logging.Record) error {
	// noop
	return nil
}

func (b *NopBackend) GetLevel(val string) logging.Level {
	return 0
}

func (b *NopBackend) SetLevel(level logging.Level, val string) {}
func (b *NopBackend) IsEnabledFor(level logging.Level, val string) bool {
	return false
}

type Logger struct {
	*logging.Logger
}

func (l Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l Logger) Println(args ...interface{}) {
	l.Info(args...)
}

func NewDummyLogger(name string) *Logger {
	gologger := logging.MustGetLogger(name)
	gologger.SetBackend(&NopBackend{})
	return &Logger{
		Logger: gologger,
	}
}

func NewLogger(name string) *Logger {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	gologger := logging.MustGetLogger(name)
	return &Logger{
		Logger: gologger,
	}
}
