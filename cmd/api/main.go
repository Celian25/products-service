package main

import (
	"context"
	"log"
	"net/http"
	"product-backend/ent"
	"product-backend/ent/migrate"
	"product-backend/internal/graphql"
	"product-backend/internal/mw"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
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

	mux := http.NewServeMux()
	srv := handler.New(graphql.NewSchema(client))
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	corsSrv := mw.CORSMiddleware(srv)
	mux.Handle("/",
		playground.Handler("Products API", "/graphql"),
	)
	mux.Handle("/graphql", corsSrv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal("http server terminated", err)
	}
}
