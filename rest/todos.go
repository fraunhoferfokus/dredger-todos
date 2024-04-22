// Edit this file, as it is a specific handler function for your service
package rest

import (
	"todos/core/log"
	"todos/core/tracing"
	"todos/usecases"
	"todos/web/pages"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// List all todo items.
func Todos(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	log.Info().Str("traceId", span.SpanContext().TraceID().String()).Str("spanId", span.SpanContext().SpanID().String()).Str("path", "/").Msg("Todos")

	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("Todos failed")
		return c.NoContent(http.StatusInternalServerError)
	}

	lzr := i18n.NewLocalizer(pages.Bundle, pages.Language(c))
	todos := usecases.Todos(session.Values["id"].(string))
	return Render(c, http.StatusOK, pages.Todos(lzr, todos))
}
