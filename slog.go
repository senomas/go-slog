package slog

import (
	"os"
	"strings"

	logrus_stack "github.com/Gurpartap/logrus-stack"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

type logger struct {
	*logrus.Logger
}

var rootLogger *logger

func init() {
	rootLogger = newRootLogger()
	logrus.AddHook(logrus_stack.NewHook(nil, []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}))
	logrus.SetLevel(logrus.InfoLevel)
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
func Debugf(format string, args ...interface{}) {
	if rootLogger.Level >= logrus.DebugLevel {
		if len(args) == 1 {
			if fx, ok := args[0].(func() interface{}); ok {
				rootLogger.Debugf(format, fx())
				return
			}
			if fx, ok := args[0].(func() []interface{}); ok {
				rootLogger.Debugf(format, fx()...)
				return
			}
		}
		rootLogger.Debugf(format, args...)
	}
}

// Infof func
func Infof(format string, args ...interface{}) {
	if rootLogger.Level >= logrus.InfoLevel {
		if len(args) == 1 {
			if fx, ok := args[0].(func() interface{}); ok {
				rootLogger.Infof(format, fx())
				return
			}
			if fx, ok := args[0].(func() []interface{}); ok {
				rootLogger.Infof(format, fx()...)
				return
			}
		}
		rootLogger.Infof(format, args...)
	}
}

// Warnf func
func Warnf(format string, args ...interface{}) {
	if len(args) == 1 {
		if fx, ok := args[0].(func() interface{}); ok {
			rootLogger.Warnf(format, fx())
			return
		}
		if fx, ok := args[0].(func() []interface{}); ok {
			rootLogger.Warnf(format, fx()...)
			return
		}
	}
	rootLogger.Warnf(format, args...)
}

// Errorf func
func Errorf(format string, args ...interface{}) {
	if len(args) == 1 {
		if fx, ok := args[0].(func() interface{}); ok {
			rootLogger.Errorf(format, fx())
			return
		}
		if fx, ok := args[0].(func() []interface{}); ok {
			rootLogger.Errorf(format, fx()...)
			return
		}
	}
	rootLogger.Errorf(format, args...)
}

// Panicf func
func Panicf(format string, args ...interface{}) {
	if len(args) == 1 {
		if fx, ok := args[0].(func() interface{}); ok {
			rootLogger.Panicf(format, fx())
			return
		}
		if fx, ok := args[0].(func() []interface{}); ok {
			rootLogger.Panicf(format, fx()...)
			return
		}
	}
	rootLogger.Panicf(format, args...)
}

// Debug func
func Debug(args ...interface{}) {
	if rootLogger.Level >= logrus.DebugLevel {
		if len(args) == 1 {
			if fx, ok := args[0].(func() interface{}); ok {
				rootLogger.Debug(fx())
				return
			}
			if fx, ok := args[0].(func() []interface{}); ok {
				rootLogger.Debug(fx()...)
				return
			}
		}
		rootLogger.Debug(args...)
	}
}

// Info func
func Info(args ...interface{}) {
	if rootLogger.Level >= logrus.InfoLevel {
		rootLogger.Info(args...)
	}
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
