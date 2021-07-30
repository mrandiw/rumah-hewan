package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Kucing struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

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

func addKucingFunc(c echo.Context) error {
	kucing := Kucing{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&kucing)
	if err != nil {
		log.Printf("Gagal melakukan decode %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status": "Gagal melakukan decode",
		})
	}

	// == save to database here ==
	log.Printf("Berhasil Menyimpan kucing dari request %v", kucing)
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Success",
	})
}

func getDashboard(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":    "Success",
		"dashboard": "yes",
	})
}

func main() {
	fmt.Println("Hello World.")

	e := echo.New()

	g := e.Group("/api/v1")

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}], host=${host}${path}, status=${status} latency=${latency}\n",
	}))

	g.GET("/dashboard", getDashboard)

	e.GET("/", home)

	e.GET("/getKucing/:type", getKucingFunc)

	e.POST("/addKucing", addKucingFunc)

	e.Start(":8080")
}
