package main

import (
	"context"
	"flag"
	"fmt"
	"gimshark-test/ui/internal/server"
	"gimshark-test/ui/pkg/config"
	"gimshark-test/ui/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

const defaultConfigFile = "conf.yaml"

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	// Get logger instance.
	log := logger.New()
	defer func() {
		done()
		log.Sync()

		if r := recover(); r != nil {
			log.Fatal("application panic", zap.Any("panic", r))
		}
	}()

	// Run server.
	err := realMain(ctx, log)

	if err != nil {
		log.Fatal("fatal err", zap.Error(err))
	}

	log.Info("successful shutdown")
}

func realMain(ctx context.Context, log *zap.Logger) error {
	// Get application cfg.
	confFlag := flag.String("conf", "", "config yaml file")
	hostFlag := flag.String("host", "", "server host")
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

	hostString := *hostFlag
	if hostString != "" {
		cfg.Packs.Host = hostString
	}

	return server.Run(ctx, log, cfg)
}
