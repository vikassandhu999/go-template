package main

import (
	"betterx/internal/controller"
	"betterx/internal/serverenv"
	"betterx/pkg/logging"
	"betterx/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-multierror"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	db postgres.Config
}

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	logger := logging.NewDefaultLogger()
	ctx = logging.NewContext(ctx, logger)

	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Error("application panic", "panic", r)
		}
	}()

	err := execMain(ctx)
	done()

	if err != nil {
		logger.Error("main failed", "error", err)
		return
	}
	logger.Info("successful shutdown")
}

func execMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	config := Config{}
	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewWithConfig(ctx, &config.db)
	if err != nil {
		return fmt.Errorf("pgxpool.newwithconfig: %w", err)
	}

	env := serverenv.New(serverenv.WithDB(db))

	router := mux.NewRouter()

	if err := controller.InitRoutes(ctx, env, router); err != nil {
		return fmt.Errorf("controller.initroutes: %w", err)
	}

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      router,
	}
	serverErrors := make(chan error, 1)

	logger.Info("starting server")

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		logger.Info("got intrupt signal, shutting down server")
		var multierr *multierror.Error

		if err := srv.Shutdown(ctx); err != nil {
			multierr = multierror.Append(err, fmt.Errorf("failed to stop server gracefully: %w", err))
		}

		if err := env.Close(ctx); err != nil {
			multierr = multierror.Append(err, fmt.Errorf("failed to close env: %w", err))
		}

		return multierr.ErrorOrNil()
	}
}
