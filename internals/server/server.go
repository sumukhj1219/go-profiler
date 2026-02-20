package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sumukhj1219/go-profiler/internals/config"
	"github.com/sumukhj1219/go-profiler/pkg/logger"
)

func Start(cfg config.Config, ctx context.Context) {
	logger.Load(cfg.Level)

	serverCtx, cancel := context.WithCancel(ctx)

	defer cancel()

	addr := cfg.Host + ":" + cfg.Port

	srv := &http.Server{
		Addr: addr,
	}

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		logger.L.Info("Shutdown signal recieved, draining connections...")
		cancel()
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.L.Fatal("server error %v", err)
		}
	}()

	logger.L.Info("server started on %s", cfg.Port)
	<-serverCtx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.L.Error("Shutdown error %v", err)
	}
	logger.L.Info("Server stopped")
}
