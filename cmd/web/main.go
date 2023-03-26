package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/samverrall/sitesmiths-api/cmd/web/api"
	"github.com/samverrall/sitesmiths-api/cmd/web/internal/repo/postgres"
	siteservice "github.com/samverrall/sitesmiths-api/internal/site"
)

var opts struct {
	http struct {
		port string
	}
}

func main() {
	flag.StringVar(&opts.http.port, "http-port", "8000", "Port for the HTTP server to listen")

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pgxConn := newDB(ctx)

	// Init repo implementations
	siteRepo := postgres.NewSiteRepo(pgxConn)

	// Init core application layer
	siteService := siteservice.New(siteRepo)

	srv := api.NewServer(api.NewServerArgs{
		Port:        opts.http.port,
		SiteService: siteService,
	})

	// Start the server on a separate Goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %s\n", err)
		}
	}()

	// Wait for a signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Shut down the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s\n", err)
	}

	// Do any cleanup or shutdown tasks
	fmt.Println("Server stopped.")
}

func newDB(ctx context.Context) *sql.DB {
	const driver = "postgres"
	db, err := sql.Open(driver, os.Getenv("DB_URL"))
	if err != nil {
		//log.Fatalf("failed to connect to postgres: %s", err.Error())
		log.Println(err.Error())
	}

	return db
}
