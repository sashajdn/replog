package log

import (
	stdlog "log"
	"sync"

	"go.uber.org/zap"
)

type Logger = zap.Logger
type SugaredLogger = zap.SugaredLogger

var (
	logger  *zap.Logger
	slogger *zap.SugaredLogger
	once    sync.Once
)

func setupLogger() {
	once.Do(func() {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			stdlog.Fatal(err)
		}

		slogger = logger.Sugar()
	})
}

func LoadSugaredLogger() *zap.SugaredLogger {
	setupLogger()
	return slogger
}

func LoadLogger() *zap.Logger {
	setupLogger()
	return logger
}
