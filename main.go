package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World Home!")
}

func getKucingFunc(c echo.Context) error {

	kucing := c.QueryParam("kucing")
	status := c.QueryParam("status")

	dataType := c.Param("type")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Ini kucingnya %s \nDan ini statusnya %s", kucing, status))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"kucing": kucing,
			"status": status,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "tipe harus string atau json",
	})
}

func main() {
	fmt.Println("Hello World.")

	e := echo.New()

	e.GET("/", home)

	e.GET("/getKucing/:type", getKucingFunc)

	e.Start(":8080")
}
