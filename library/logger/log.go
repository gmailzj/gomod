package logger

import (
	"github.com/sirupsen/logrus"
	ft "gomod/library/logger/formatter"
	"gomod/library/logger/hooks"
	"io"
	"reflect"
	"sync"
)

var logger *logrus.Entry

var formatter logrus.Formatter

func InitLogger(app string, writer io.Writer, args ...interface{}) {
	var once sync.Once
	once.Do(func() {
		var format logrus.Formatter
		var chks CustomHooks
		var customWithFields CustomFields
		var logLevel logrus.Level = 999

		for _, v := range args {
			if f, ok := v.(logrus.Formatter); ok {
				format = f
			}
			if h, ok := v.(CustomHooks); ok {
				chks = h
			}
			if fd, ok := v.(CustomFields); ok {
				customWithFields = fd
			}
			if l, ok := v.(logrus.Level); ok {
				logLevel = l
			}
		}

		hks := defaultHooks()
		for fn, hook := range chks.Hooks {
			if k, ok := hook.(bool); ok {
				if k == false { //取消默认的钩子
					delete(hks, fn)
				}
			}
			if k, ok := hook.(logrus.Hook); ok {
				hks[fn] = k //追加自定义钩子
			}
		}

		withFields := defaultFields(app)
		for fn, field := range customWithFields.WithFields {
			if k, ok := field.(bool); ok {
				if k == false { //取消默认field
					delete(withFields, fn)
				}
			}
			if k, ok := field.(map[string]string); ok {
				withFields[fn] = k //追加自定义field
			}
		}

		if logLevel == 999 {
			logLevel = defaultLevel()
		}

		if formatter == nil {
			format = &ft.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05", DataKey: "context", FieldMap: ft.FieldMap{
				logrus.FieldKeyTime:  "logTime",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "@caller",
			},
			}
		}

		formatter = format

		log := logrus.New()
		log.Out = writer
		log.Formatter = formatter
		log.Level = logLevel

		for _, hook := range hks {
			log.Hooks.Add(hook)
		}

		logger = log.WithFields(withFields)
	})
}

func defaultFields(app string) logrus.Fields {
	return logrus.Fields{"app": app}
}

func defaultHooks() map[string]logrus.Hook {
	f := hooks.NewFormatLogIdHook()
	l := hooks.NewLineHook()
	h := map[string]logrus.Hook{
		reflect.TypeOf(f).String(): f,
		reflect.TypeOf(l).String(): l,
	}
	if hooks.EnableSentry {
		s := hooks.NewSentryHook()
		h[reflect.TypeOf(s).String()] = s
	}
	return h
}

func defaultLevel() logrus.Level {
	return logrus.InfoLevel
}

func GetLogger() *logrus.Entry {
	return logger
}

func Trace(args ...interface{}) {
	logger.Trace(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Print(args ...interface{}) {
	logger.Print(args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

func Panic(args ...interface{}) {
	logger.Panic(args)
}

// Entry Printf family functions

func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

func Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}

// Entry Println family functions

func Traceln(args ...interface{}) {
	logger.Traceln(args)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

func Println(args ...interface{}) {
	logger.Println(args)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

func Warningln(args ...interface{}) {
	logger.Warningln(args)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args)
}
