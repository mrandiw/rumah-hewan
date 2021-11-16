package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "template",
		Index: "index.html",
	}))

	// Server header
	e.Use(serverHeader)
}

// === MIDDLEWARE ===
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Rumah Hewan 1.1")
		c.Response().Header().Set("Developer", "Andi Wibowo")
		return next(c)
	}
}
