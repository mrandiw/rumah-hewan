package handlers

import (
	"log"
	"net/http"
	"rumah-hewan/db"
	"time"

	"github.com/labstack/echo"
)

type Animal struct {
	ID        uint
	Name      string
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Home(c echo.Context) error {
	if err := db.Open(); err != nil {
		log.Fatal("koneksi database gagal", err)
	}
	defer db.Close()

	var animal []Animal
	result := db.DB.Find(&animal)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, map[string][]Animal{
		"animal": animal,
	})
}
