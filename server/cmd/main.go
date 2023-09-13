package main

import (
	"context"
	"flag"
	"fmt"
	"gimshark-test/server/internal/server"
	"gimshark-test/server/pkg/config"
	"gimshark-test/server/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

const defaultConfigFile = "conf.yaml"

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	// Get logger interface.
	log := logger.New()
	defer func() {
		done()
		log.Sync()

		if r := recover(); r != nil {
			fmt.Println(1)
			os.Exit(1)
		}
	}()

	// Run server.
	err := realMain(ctx, log)

	if err != nil {
		log.Error("fatal err", zap.Error(err))
		panic(err)
	}

	log.Info("successful shutdown")
}

func realMain(ctx context.Context, log *zap.Logger) error {
	// Get application cfg.
	confFlag := flag.String("conf", "", "cfg. yaml file")
	flag.Parse()

	confString := *confFlag
	if confString == "" {
		confString = defaultConfigFile
	}

	cfg, err := config.Parse(confString)
	if err != nil {
		log.Error("failed to parse config", zap.Error(err))

		return fmt.Errorf("failed to parse config: %w", err)
	}

	return server.Run(ctx, log, cfg)
}
