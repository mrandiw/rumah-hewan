package router

import (
	"rumah-hewan/api"
	"rumah-hewan/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {

	e := echo.New()

	// create group
	g := e.Group("/api/v1")
	gCookie := e.Group("/cookie/v1")
	gJwt := e.Group("/jwt/v1")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddleware(g)
	middlewares.SetCookieMiddlewares(gCookie)
	middlewares.SetJwtMiddlewares(gJwt)

	// set main router
	api.MainsGroup(e)

	// set group router
	api.AdminGroup(g)
	api.CookieGroup(gCookie)
	api.JwtGroup(gJwt)

	return e
}
