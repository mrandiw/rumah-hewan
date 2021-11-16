package api

import (
	"rumah-hewan/api/handlers"

	"github.com/labstack/echo"
)

func CookieGroup(e *echo.Group) {
	e.GET("/main", handlers.GetDashboardCookie)

}
