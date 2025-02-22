package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("failed to create logger", zap.Error(err))
	}

	slogger := logger.Sugar()
	slogger.Infof(`Starting replog application...`)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	<-ctx.Done()
	slogger.Infof("Replog application closing down, shutdown signal received")
}
