package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"go.uber.org/zap"

	"github.com/sashajdn/replog/pkg/application/replog"
	"github.com/sashajdn/replog/pkg/connectivity/discord"
	"github.com/sashajdn/replog/pkg/connectivity/repositories"
	"github.com/sashajdn/replog/pkg/db"
	"github.com/sashajdn/replog/pkg/env"
)

func main() {
	// TODO: prod vs dev
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(`Failed to create logger`, zap.Error(err))
	}

	slogger := logger.Sugar()
	slogger.Infof("Starting replog application...")

	env, err := env.Load()
	if err != nil {
		slogger.With(zap.Error(err)).Error("Failed to load env")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	db, err := db.New(env.PostgreSQL)
	if err != nil {
		slogger.With(zap.Error(err)).Fatalf("Failed to create db instance")
	}

	discordClient, err := discord.NewClient(discord.ClientConfig{
		Logger:              slogger,
		ReceiverChannelSize: 100_000,
		Token:               env.Discord.APIToken,
	})
	if err != nil {
		slogger.With(zap.Error(err)).Fatal("Failed to create discord client")
	}

	entryRepo := repositories.NewEntryRepository(slogger, db)
	userRepo := repositories.NewUserRepository(slogger, db)
	service := replog.NewService(slogger, discordClient, entryRepo, userRepo)

	if err := service.Run(ctx); err != nil {
		slogger.With(zap.Error(err)).Fatal("Failed to run repolog service")
	}

	<-ctx.Done()
	slogger.Infof("Replog application closing down, shutdown signal received")
}
