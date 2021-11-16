package api

import (
	"rumah-hewan/api/handlers"

	"github.com/labstack/echo"
)

func JwtGroup(e *echo.Group) {
	e.GET("/main", handlers.GetDashboardJwt)

}
