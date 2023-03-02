package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"tinderclone_back/src/pkg/domain"
)

type DbConfig struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

var database *gorm.DB

func InitializeDb(config *DbConfig) {
	var err error
	dsn := "host=" + config.DbHost + " user=" + config.DbUsername + " dbname=" + config.DbName + " port=" + config.DbPort
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to db")
	}
	err = database.AutoMigrate(&domain.User{})
	if err != nil {
		panic("Couldn't migrate")
	}
}

func GetDb() *gorm.DB {
	return database
}
