// Edit this file, as it is a specific handler function for your service
package rest

import (
	"dredgerTodos/core/log"
	"dredgerTodos/core/tracing"
	"dredgerTodos/usecases/messages"
	"dredgerTodos/web/pages"

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

	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Str("session", session.Values["id"].(string)).Msg("Todos failed")
		return c.NoContent(http.StatusInternalServerError)
	}

	log.Info().Str("session", session.Values["id"].(string)).Str("traceId", span.SpanContext().TraceID().String()).Str("spanId", span.SpanContext().SpanID().String()).Str("path", "/").Msg("Todos")

	lzr := i18n.NewLocalizer(pages.Bundle, pages.Language(c))
	todos := messages.Todos(session.Values["id"].(string))
	return Render(c, http.StatusOK, pages.Todos(lzr, todos))
}
