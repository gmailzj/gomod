package logger

import (
	"github.com/sirupsen/logrus"
	baseLog "gomod/library/logger"
	"io"
	"os"
	"sync"
)

var onceInit sync.Once

var log *logrus.Entry

var logPath = ""

var logName = "application.log"

var hooks baseLog.CustomHooks

var withFields baseLog.CustomFields

var writer io.Writer = nil

var logLevel logrus.Level = logrus.InfoLevel

func SetLogPath(p string) {
	logPath = p
}

func SetLogName(l string) {
	logName = l
}

func SetHooks(h baseLog.CustomHooks) {
	hooks = h
}

func SetWithFields(f baseLog.CustomFields) {
	withFields = f
}

func SetWriter(w io.Writer) {
	writer = w
}

func SetLogLevel(l uint32) {
	logLevel = logrus.Level(l)
}

func InitLogger() {
	onceInit.Do(func() {
		app := "log"
		str, _ := os.Getwd()
		if logPath == "" {
			logPath = str + "/runtime/"
			err := os.MkdirAll(logPath, os.ModePerm)
			if err != nil {
				panic("创建日志目录错误" + err.Error())
				return
			}
		}
		logFile := logPath + logName
		var w io.Writer
		if writer == nil {
			var err error
			w, err = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil && os.IsNotExist(err) {
				_, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
				if err != nil {
					panic("创建日志文件错误" + err.Error())
				}
				w, err = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			}
			if err != nil {
				panic("打开日志文件错误" + err.Error())
			}
		} else {
			w = writer
		}
		baseLog.InitLogger(app, w, withFields, hooks, logLevel)
		log = baseLog.GetLogger()
	})

}

func Trace(args ...interface{}) {
	log.Trace(args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Print(args ...interface{}) {
	log.Print(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Warning(args ...interface{}) {
	log.Warning(args...)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Panic(args ...interface{}) {
	log.Panic(args)
}

// Entry Printf family functions

func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args)
}

func Warningf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args)
}

// Entry Println family functions

func Traceln(args ...interface{}) {
	log.Traceln(args)
}

func Debugln(args ...interface{}) {
	log.Debugln(args)
}

func Infoln(args ...interface{}) {
	log.Infoln(args)
}

func Println(args ...interface{}) {
	log.Println(args)
}

func Warnln(args ...interface{}) {
	log.Warnln(args)
}

func Warningln(args ...interface{}) {
	log.Warningln(args)
}

func Errorln(args ...interface{}) {
	log.Errorln(args)
}

func Fatalln(args ...interface{}) {
	log.Fatalln(args)
}

func Panicln(args ...interface{}) {
	log.Panicln(args)
}

func GetLogger() *logrus.Entry {
	//InitLogger()
	return log
}
