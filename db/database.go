package db

import (
	"log"

	"githu.com/prabhumohan2000/test_8/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

var err error

func Init() {
	dbURL := "postgres://postgres:postgres@localhost:5432/test_db"
	DbInstance, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("error occured in connecting postgres database")
	}
	DbInstance.AutoMigrate(&model.User{})
}
