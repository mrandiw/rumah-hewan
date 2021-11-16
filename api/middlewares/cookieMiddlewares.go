package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("SessionName")
		if err != nil {
			return c.String(http.StatusUnauthorized, "Authentication Failed")
		}

		if cookie.Value == "SessionValue" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Authentication Failed")
	}
}
