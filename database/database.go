package database

import (
	"booooooook/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=BookCRUDTEST port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Book{})

	return db
}
