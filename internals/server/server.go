package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sumukhj1219/go-profiler/internals/config"
	"github.com/sumukhj1219/go-profiler/internals/handler"
	"github.com/sumukhj1219/go-profiler/internals/repository"
	"github.com/sumukhj1219/go-profiler/internals/routes"
	"github.com/sumukhj1219/go-profiler/internals/service"
	"github.com/sumukhj1219/go-profiler/pkg/database"
	"github.com/sumukhj1219/go-profiler/pkg/logger"
)

func Start(cfg config.Config, ctx context.Context) {
	logger.Load(cfg.Log.Level)

	db, err := database.New(ctx, cfg)
	if err != nil {
		logger.L.Fatal("Cannot connect to database %v", err)
	}
	defer db.Close()

	router := mux.NewRouter()
	serverCtx, cancel := context.WithCancel(ctx)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewService(userRepo)

	handlers := handler.Handler{
		UserHandler: handler.NewHandler(userService, *logger.L),
	}

	routes.RegisterRoutes(router, handlers)

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
