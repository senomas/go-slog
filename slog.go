package slog

import (
	"os"
	"strings"

	logrus_stack "github.com/Gurpartap/logrus-stack"
	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

type logger struct {
	*logrus.Logger
}

var (
	rootLogger = newRootLogger()
)

func init() {
	logrus.AddHook(logrus_stack.NewHook(nil, []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}))
}

func newRootLogger() *logger {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(newFormatter(true))
	return &logger{logrus.StandardLogger()}
}

// SetLevelToDebug func
func SetLevelToDebug() {
	logrus.SetLevel(logrus.DebugLevel)
}

// SetLevelToInfo func
func SetLevelToInfo() {
	logrus.SetLevel(logrus.InfoLevel)
}

// SetLevelToWarn func
func SetLevelToWarn() {
	logrus.SetLevel(logrus.WarnLevel)
}

// Debugf func
func Debugf(format string, fn func() interface{}) {
	if rootLogger.Level >= logrus.DebugLevel {
		ff := fn()
		rootLogger.Debugf(format, ff)
	}
}

// Debugff func
func Debugff(format string, fn func() []interface{}) {
	if rootLogger.Level >= logrus.DebugLevel {
		ff := fn()
		rootLogger.Debugf(format, ff...)
	}
}

// Debugfi func
func Debugfi(format string, ff ...interface{}) {
	if rootLogger.Level >= logrus.DebugLevel {
		rootLogger.Debugf(format, ff...)
	}
}

// Infof func
func Infof(format string, fn func() interface{}) {
	if rootLogger.Level >= logrus.InfoLevel {
		ff := fn()
		rootLogger.Infof(format, ff)
	}
}

// Infoff func
func Infoff(format string, fn func() []interface{}) {
	if rootLogger.Level >= logrus.InfoLevel {
		ff := fn()
		rootLogger.Infof(format, ff)
	}
}

// Infofi func
func Infofi(format string, ff ...interface{}) {
	if rootLogger.Level >= logrus.DebugLevel {
		rootLogger.Infof(format, ff...)
	}
}

// Warnf func
func Warnf(format string, args ...interface{}) {
	rootLogger.Warnf(format, args...)
}

// Errorf func
func Errorf(format string, args ...interface{}) {
	rootLogger.Errorf(format, args...)
}

// Panicf func
func Panicf(format string, args ...interface{}) {
	rootLogger.Panicf(format, args...)
}

// Debug func
func Debug(args ...interface{}) {
	if rootLogger.Level >= logrus.DebugLevel {
		rootLogger.Debug(args...)
	}
}

// Info func
func Info(args ...interface{}) {
	rootLogger.Info(args...)
}

// Warn func
func Warn(args ...interface{}) {
	rootLogger.Warn(args...)
}

// Error func
func Error(args ...interface{}) {
	rootLogger.Error(args...)
}

// Panic func
func Panic(args ...interface{}) {
	rootLogger.Panic(args...)
}
