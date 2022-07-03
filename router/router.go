package router

import (
	"rumah-hewan/api/handlers"
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
	MainRouter(e)

	// set group router
	AdminRouter(g)
	CookieRouter(gCookie)
	JwtRouter(gJwt)

	return e
}

func MainRouter(e *echo.Echo) {
	e.GET("/home", handlers.Home)
	e.GET("/login", handlers.Login)

	e.GET("/getKucing/:type", handlers.GetKucingFunc)

	e.POST("/addKucing", handlers.AddKucingFunc)
}

func AdminRouter(g *echo.Group) {
	g.GET("/dashboard", handlers.GetDashboard)
}

func CookieRouter(e *echo.Group) {
	e.GET("/main", handlers.GetDashboardCookie)

}

func JwtRouter(e *echo.Group) {
	e.GET("/main", handlers.GetDashboardJwt)

}
