package db

import (
	"rumah-hewan/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Open() error {
	var err error

	config, _ := config.LoadConfig(".")

	dsn := config.Database.Mysql

	DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		return err
	}

	return err
}

func Close() error {
	return DB.Close()
}
