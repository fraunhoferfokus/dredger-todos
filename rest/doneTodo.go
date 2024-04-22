// Edit this file, as it is a specific handler function for your service
package rest

import (
	"todos/core/log"
	"todos/core/tracing"
	"todos/usecases"
	"todos/web/pages"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Mark todo item as Done.
func DoneTodo(c echo.Context) error {
	// trace span
	ctx := c.Request().Context()
	ctx, span := tracing.Tracer.Start(ctx, "logMessage")
	defer span.End()

	log.Info().Str("traceId", span.SpanContext().TraceID().String()).Str("spanId", span.SpanContext().SpanID().String()).Str("path", "/").Msg("DoneTodo")

	session, err := getSession(c)
	if err != nil {
		log.Error().Err(err).Msg("DoneTodo failed")
		return c.NoContent(http.StatusInternalServerError)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	ses := session.Values["id"].(string)

	usecases.DoneTodo(ses, id)
	lzr := i18n.NewLocalizer(pages.Bundle, pages.Language(c))
	todos := usecases.Todos(ses)
	return Render(c, http.StatusOK, pages.Todos(lzr, todos))
}
