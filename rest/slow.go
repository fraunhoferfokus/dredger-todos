// Edit this file, as it is a specific handler function for your service
package rest

import (
	"dredgerTodos/core/log"
	"dredgerTodos/core/tracing"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

// service answer slow
func Slow(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	log.Info().Str("traceId", span.SpanContext().TraceID().String()).Str("spanId", span.SpanContext().SpanID().String()).Str("path", "/").Msg("Slow")

	// implement your functionality best using a function from a separate file, e.g. SlowDo.go
	time.Sleep(time.Second * 15)

	// 200 => deliver slow answer (10 sec) for testing load bars
	return c.String(http.StatusOK, "I'm slow.")
}
