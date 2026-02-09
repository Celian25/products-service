package graphql

import (
	"log/slog"
	"net/http"
	"product-backend/ent"
	"product-backend/internal/mw"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func New(db *ent.Client, logger *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	srv := handler.New(NewSchema(db))
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	srv = mw.LoggerMiddleware(srv, logger)

	mwSrv := mw.CORSMiddleware(srv)
	mux.Handle("/",
		playground.Handler("Products API", "/graphql"),
	)
	mux.Handle("/graphql", mwSrv)

	return mux
}
