package main

import (
	"context"
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
	"github.com/samverrall/sitesmiths-api/cmd/internal/repo/google"
	"github.com/samverrall/sitesmiths-api/cmd/internal/repo/mongodb"
	"github.com/samverrall/sitesmiths-api/cmd/web/api"
	"github.com/samverrall/sitesmiths-api/internal/account"
	"github.com/samverrall/sitesmiths-api/internal/site"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var opts struct {
	http struct {
		port     string
		insecure bool
	}
}

const (
	sitesCollection = "sites"
)

func main() {
	flag.StringVar(&opts.http.port, "http-port", "8000", "Port for the HTTP server to listen")

	// TODO: Parse config

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set client options

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_DB_URL")).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to mongo: %s", err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("failed to ping mongo: %s", err.Error())
	}

	database := client.Database(os.Getenv("MONGO_DB_NAME"))

	// Init repo implementations
	siteRepo := mongodb.NewSiteRepo(database.Collection(sitesCollection))
	accountRepo := mongodb.NewAccountRepo(database)

	authenticatorRepo := google.NewAuthenticator("TODO", "TODO", "TODO")

	// Init core application layer
	siteService := site.NewService(siteRepo)
	accountService := account.NewService(accountRepo, authenticatorRepo)

	api := api.New(siteService, accountService, opts.http.port, opts.http.insecure)

	srv := api.NewServer()

	// Start the server on a separate Goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %s\n", err.Error())
		}
	}()

	// Wait for a signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Shut down the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s\n", err.Error())
	}

	// Do any cleanup or shutdown tasks
	fmt.Println("Server stopped.")
}
