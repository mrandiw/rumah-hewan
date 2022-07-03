package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rumah-hewan/db"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

type Animal struct {
	ID        uint
	Name      string
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// home is public : first function name capital
func Home(c echo.Context) error {
	if err := db.Open(); err != nil {
		log.Fatal("koneksi database gagal", err)
	}
	defer db.Close()

	animal, err := getAnimal()
	if err != nil {
		return err
	}

	// mongo connection
	client := db.MgoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return err
	}
	fmt.Println(databases)

	return c.JSON(http.StatusOK, map[string][]Animal{
		"animal": animal,
	})
}

// getAnimal is privive : first function name not capital
func getAnimal() ([]Animal, error) {

	var animal []Animal
	result := db.DB.Find(&animal)
	if result.Error != nil {
		return animal, result.Error
	}

	return animal, nil
}
