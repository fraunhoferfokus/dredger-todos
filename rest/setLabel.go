// Edit this file, as it is a specific handler function for your service
package rest

import (
	"dredgerTodos/core/log"
	"dredgerTodos/core/tracing"
	"dredgerTodos/usecases"
	"dredgerTodos/usecases/messages"
	"dredgerTodos/web/pages"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Label the todo set.
func SetLabel(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	spanId := span.SpanContext().SpanID().String()
	log.Info().Str("traceId", traceId).Str("spanId", spanId).Str("path", "/").Msg("SetLabel")

	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("SetLabel failed")
		return c.NoContent(http.StatusInternalServerError)
	}

	label := c.FormValue("label")
	usecases.CurrentLabels.Set(label, session.Values["id"].(string))

	log.Info().Str("session", session.Values["id"].(string)).Str("traceId", span.SpanContext().TraceID().String()).Str("spanId", span.SpanContext().SpanID().String()).Str("label", label).Str("path", "/").Msg("SetLabel")

	lzr := i18n.NewLocalizer(pages.Bundle, pages.Language(c))
	todos := messages.Todos(session.Values["id"].(string))
	return Render(c, http.StatusOK, pages.Todos(lzr, todos))
}
