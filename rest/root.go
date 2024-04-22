// Edit this file, as it is a specific handler function for your service
package rest

import (
	"todos/core/log"
	"todos/core/tracing"
	"todos/web/pages"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Returns the index page
func Root(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	log.Info().Str("traceId", span.SpanContext().TraceID().String()).Str("spanId", span.SpanContext().SpanID().String()).Str("path", "/").Msg("Root")

	lzr := i18n.NewLocalizer(pages.Bundle, pages.Language(c))
	return Render(c, http.StatusOK, pages.Index(lzr))
}
