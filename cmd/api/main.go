package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"product-backend/ent"
	"product-backend/ent/migrate"
	"product-backend/internal/graphql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "./products.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed migrating schema: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := graphql.New(client, logger)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		logger.Info("server started", "port", 8081)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("http server terminated", err)
		}
	}()
	<-stop

	logger.Error("server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	logger.Error("server exited")
}
