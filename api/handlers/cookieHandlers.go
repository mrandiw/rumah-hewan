package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetDashboardCookie(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Success",
		"page":   "cookie",
	})
}
