package api

import (
	"rumah-hewan/api/handlers"

	"github.com/labstack/echo"
)

func MainsGroup(e *echo.Echo) {
	e.GET("/home", handlers.Home)
	e.GET("/login", handlers.Login)

	e.GET("/getKucing/:type", handlers.GetKucingFunc)

	e.POST("/addKucing", handlers.AddKucingFunc)

}
