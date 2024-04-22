// Edit this file, as it is a specific handler function for your service
package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Returns ready-state of the service
func GetReady(c echo.Context) error {
	return c.String(http.StatusOK, "I'm ready.")
}
