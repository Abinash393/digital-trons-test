package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ORM *gorm.DB

func Initialize() {
	dsn := "host=localhost user=root password=password dbname=trons port=5432 sslmode=disable"
	var err error
	ORM, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
