package log

import (
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Sirupsen/logrus"
)

/*************************************************
Debug Level Setting
- debug
- info
- warning
- error
- fatal
- panic
*************************************************/

const (
	fileTag = "file"
	lineTag = "line"
	funcTag = "func"
)

type logConf struct { //執行時期的 log 功能配置
	showFileInfo bool //是否顯示 file name, func name, line number
}

var rtLogConf logConf

//InitLog config the log
func InitLog(level logrus.Level, filename string, MultiWriter bool, showFileInfo bool) {
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: false, FullTimestamp: true})
	err := os.MkdirAll(filepath.Dir(filename), 0744)
	if err != nil {
		Error("error folder create : ", err)
		os.Exit(1)
	}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		Error("error opening file: ", err)
		os.Exit(1)
	}
	if MultiWriter {
		logrus.SetOutput(io.MultiWriter(f, os.Stdout))
	} else {
		logrus.SetOutput(f)
	}
	rtLogConf.showFileInfo = showFileInfo

}

func Init(environment string, filename string, level string) {
	debugLV, _ := logrus.ParseLevel(level)
	switch environment {
	case "development":
		InitLog(debugLV, filename, true, true)
	case "production":
		InitLog(debugLV, filename, false, false)
	}
}

func getBaseName(fileName string, funcName string) (string, string) {
	return filepath.Base(fileName), filepath.Base(funcName)
}

// Println same as Debug
func Println(args ...interface{}) {
	Debug(args)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if !rtLogConf.showFileInfo {
		logrus.Debug(args...)
		return
	}

	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Debug(args...)
	} else {
		logrus.Debug(args...)
	}
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(msg string, args ...interface{}) {
	if !rtLogConf.showFileInfo {
		logrus.Debugf(msg, args...)
		return
	}

	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Debugf(msg, args...)
	} else {
		logrus.Debugf(msg, args...)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if !rtLogConf.showFileInfo {
		logrus.Info(args...)
		return
	}

	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Info(args...)
	} else {
		logrus.Info(args...)
	}
}

// Infof logs a message at level Info on the standard logger.
func Infof(msg string, args ...interface{}) {
	if !rtLogConf.showFileInfo {
		logrus.Infof(msg, args...)
		return
	}

	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Infof(msg, args...)
	} else {
		logrus.Infof(msg, args...)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(msg ...interface{}) {
	if !rtLogConf.showFileInfo {
		logrus.Warn(msg...)
		return
	}

	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Warn(msg...)
	} else {
		logrus.Warn(msg...)
	}
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(msg string, args ...interface{}) {
	if !rtLogConf.showFileInfo {
		logrus.Warnf(msg, args...)
		return
	}

	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Warnf(msg, args...)
	} else {
		logrus.Warnf(msg, args...)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(msg ...interface{}) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Error(msg...)
	} else {
		logrus.Error(msg...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(msg string, args ...interface{}) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Errorf(msg, args...)
	} else {
		logrus.Errorf(msg, args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(msg ...interface{}) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Fatal(msg...)
	} else {
		logrus.Fatal(msg...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(msg string, args ...interface{}) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Fatalf(msg, args...)
	} else {
		logrus.Fatalf(msg, args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(msg ...interface{}) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Panic(msg...)
	} else {
		logrus.Panic(msg...)
	}
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(msg string, args ...interface{}) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fileName, funcName := getBaseName(file, runtime.FuncForPC(pc).Name())
		logrus.WithField(fileTag, fileName).WithField(lineTag, line).WithField(funcTag, funcName).Panicf(msg, args...)
	} else {
		logrus.Panicf(msg, args...)
	}
}
