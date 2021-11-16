package api

import (
	"rumah-hewan/api/handlers"

	"github.com/labstack/echo"
)

func AdminGroup(g *echo.Group) {
	g.GET("/dashboard", handlers.GetDashboard)
}
