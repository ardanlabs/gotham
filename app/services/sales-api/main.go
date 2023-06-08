package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ardanlabs/service/foundation/logger"
)

func main() {
	log := logger.New(os.Stdout, "SALES-API")

	if err := run(log); err != nil {
		log.Info("startup", "ERROR", err)
		os.Exit(1)
	}
}

func run(log *logger.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	sig := <-shutdown

	log.Info("shutdown", "status", "shutdown started", "signal", sig)
	defer log.Info("shutdown", "status", "shutdown complete", "signal", sig)

	return nil
}
