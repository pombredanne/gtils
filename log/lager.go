package log

import (
	"flag"
	"fmt"
	// "io"

	"github.com/pivotal-golang/lager"
)

func init() {
	AddFlags(flag.CommandLine)
	flag.Parse()
}

func NewLager(log *logger) Logger {
	var minLagerLogLevel lager.LogLevel
	switch minLogLevel {
	case DEBUG:
		minLagerLogLevel = lager.DEBUG
	case INFO:
		minLagerLogLevel = lager.INFO
	case ERROR:
		minLagerLogLevel = lager.ERROR
	case FATAL:
		minLagerLogLevel = lager.FATAL
	default:
		panic(fmt.Errorf("unknown log level: %s", minLogLevel))
	}

	logger := lager.NewLogger(log.Name)
	logger.RegisterSink(lager.NewWriterSink(log.Writer, minLagerLogLevel))
	log.Logger = logger

	return log
}

// func (l *logger) GetSink() io.Writer {
// 	return l.Writer
// }

func (l *logger) Debug(action string, data ...Data) {
	l.Logger.Debug(action, toLagerData(data...))
}

func (l *logger) Info(action string, data ...Data) {
	l.Logger.Info(action, toLagerData(data...))
}

func (l *logger) Error(action string, err error, data ...Data) {
	l.Logger.Error(action, err, toLagerData(data...))
}

func (l *logger) Fatal(action string, err error, data ...Data) {
	l.Logger.Fatal(action, err, toLagerData(data...))
}

func toLagerData(givenData ...Data) lager.Data {
	data := lager.Data{}

	if len(givenData) > 0 {
		for _, dataArg := range givenData {
			for key, val := range dataArg {
				data[key] = val
			}
		}
	}

	return data
}
