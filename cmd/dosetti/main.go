package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/golang-cz/devslog"
	"github.com/jariinc/dosetti/internal/database"
	"github.com/jariinc/dosetti/internal/database/libsql"
	"github.com/jariinc/dosetti/internal/server"
	"github.com/joho/godotenv"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	logger := slog.New(devslog.NewHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		slog.Error(".env loading failed", "error", err)
	}

	db, err := libsql.NewConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	defer db.Close()

	// Migrate the database before starting the server
	slog.Info("migrating database...")
	if err := database.Migrate(ctx, db); err != nil {
		panic(err)
	}
	slog.Info("migration done")

	server := server.NewServer(db)

	httpServer := &http.Server{
		Addr:         net.JoinHostPort("", "8080"),
		Handler:      server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	go func() {
		slog.Info("listening TPC portg", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error listening and serving", "error", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Error("error shutting down http server", "error", err)
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		slog.Error("main exiting", "error", err)
		os.Exit(1)
	}
}
