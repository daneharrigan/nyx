package logger

import (
	"io"
	"log"
	"os"
)

type Logger interface {
	// builtin
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
	SetPrefix(string)
	// info
	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})
	// debug
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
	// warn
	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})
	// error
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})
	// context
	WithContext(string, func(Logger))
	// level
	SetLevel(Level)
}

type logger struct {
	*log.Logger
	out   io.Writer
	level Level
}

type Level int

const (
	LevelInfo Level = iota
	LevelDebug
	LevelWarn
	LevelError
)

var std Logger = New(os.Stderr, "")

func Info(v ...interface{}) {
	std.Info(v...)
}

func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

func Warn(v ...interface{}) {
	std.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

func Warnln(v ...interface{}) {
	std.Warnln(v...)
}

func Error(v ...interface{}) {
	std.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

func WithContext(prefix string, fn func(Logger)) {
	std.WithContext(prefix, fn)
}

func SetLevel(n Level) {
	std.SetLevel(n)
}

func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}

func Panic(v ...interface{}) {
	std.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	std.Panicln(v...)
}

func Print(v ...interface{}) {
	std.Print(v...)
}

func Printf(format string, v ...interface{}) {
	std.Printf(format, v...)
}

func Println(v ...interface{}) {
	std.Println(v...)
}

func New(out io.Writer, prefix string) Logger {
	if out == nil {
		out = os.Stderr
	}

	l := &logger{log.New(out, "", 0), out, LevelInfo}
	l.SetPrefix(prefix)

	return l
}

func (l *logger) Info(v ...interface{}) {
	print(l, LevelInfo, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	printf(l, LevelInfo, format, v...)
}

func (l *logger) Infoln(v ...interface{}) {
	println(l, LevelInfo, v...)
}

func (l *logger) Debug(v ...interface{}) {
	print(l, LevelDebug, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	printf(l, LevelDebug, format, v...)
}

func (l *logger) Debugln(v ...interface{}) {
	println(l, LevelDebug, v...)
}

func (l *logger) Warn(v ...interface{}) {
	print(l, LevelWarn, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	printf(l, LevelWarn, format, v...)
}

func (l *logger) Warnln(v ...interface{}) {
	println(l, LevelWarn, v...)
}

func (l *logger) Error(v ...interface{}) {
	print(l, LevelError, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	printf(l, LevelError, format, v...)
}

func (l *logger) Errorln(v ...interface{}) {
	println(l, LevelError, v...)
}

func (l *logger) WithContext(prefix string, fn func(Logger)) {
	if l.Prefix() != "" {
		prefix = l.Prefix() + " " + prefix
	}

	ctx := New(l.out, prefix)
	fn(ctx)
}

func (l *logger) SetLevel(n Level) {
	l.level = n
}

func (l *logger) SetPrefix(prefix string) {
	if prefix != "" {
		prefix += " "
	}

	l.Logger.SetPrefix(prefix)
}

func print(l *logger, level Level, v ...interface{}) {
	if level >= l.level {
		l.Print(v...)
	}
}

func printf(l *logger, level Level, format string, v ...interface{}) {
	if level >= l.level {
		l.Printf(format, v...)
	}
}

func println(l *logger, level Level, v ...interface{}) {
	if level >= l.level {
		l.Println(v...)
	}
}
