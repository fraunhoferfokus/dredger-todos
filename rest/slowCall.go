// Edit this file, as it is a specific handler function for your service
package rest

import (
	"dredgerTodos/core/log"
	"dredgerTodos/core/tracing"

	"net/http"

	"github.com/labstack/echo/v4"
)

// service call slow
func SlowCall(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	log.Info().Str("traceId", span.SpanContext().TraceID().String()).Str("spanId", span.SpanContext().SpanID().String()).Str("path", "/").Msg("SlowCall")

	f := func() {
		// implement your functionality best using a function from a separate file, e.g. usecases/SlowCallDo.go
		_, err := http.Get("http://localhost:9090/slowz")
		if err != nil {
			log.Warn().Err(err).Msg("Slow call failed")
		}
	}
	ProgressPico(f)

	// 200 => call slow for testing load bars
	return c.String(http.StatusOK, "I'm slow.")
}
