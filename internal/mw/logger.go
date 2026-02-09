package mw

import (
	"context"
	"log/slog"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
)

func LoggerMiddleware(srv *handler.Server, logger *slog.Logger) *handler.Server {
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		start := time.Now()

		responseHandler := next(ctx)
		if responseHandler == nil {
			logger.Warn("graphql operation returned nil response handler")
			return nil
		}

		var logged bool

		return func(ctx context.Context) *graphql.Response {
			opCtx := graphql.GetOperationContext(ctx)

			if !logged {
				logger.Info(
					"graphql operation started",
					slog.String("operation", opCtx.OperationName),
					slog.String("operation_type", string(opCtx.Operation.Operation)),
				)
				logged = true
			}

			resp := responseHandler(ctx)

			if resp == nil {
				logger.Info(
					"graphql operation finished",
					slog.String("operation", opCtx.OperationName),
					slog.Duration("duration", time.Since(start)),
				)
				return nil
			}

			if len(resp.Errors) > 0 {
				logger.Error(
					"graphql operation error",
					slog.String("operation", opCtx.OperationName),
					slog.Int("errors_count", len(resp.Errors)),
					slog.Any("errors", resp.Errors),
				)
			}

			return resp
		}
	})

	return srv
}
