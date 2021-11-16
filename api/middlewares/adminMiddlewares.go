package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddleware(g *echo.Group) {
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}], host=${host}${path}, status=${status} latency=${latency}\n",
	}))

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {

		// query ke database

		// Be careful to use constant time comparison to prevent timing attacks
		if username == "andi" && password == "123456" {
			return true, nil
		}
		return false, nil
	}))
}
