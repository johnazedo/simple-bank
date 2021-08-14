package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"news/domain"
)

var db *gorm.DB = nil

func DatabaseSetup(){
	var err error;
	db, err = gorm.Open(sqlite.Open("news.db"), &gorm.Config{});

	if err != nil {
		panic("Failed to connect database")
	}

	MigrateData(db, &domain.News{})
}

func GetDB() *gorm.DB {
	return db
}

func MigrateData(db *gorm.DB, schema interface{})  {
	err := db.AutoMigrate(schema)
	if err != nil {
		panic("Cannot migrate schema: " + fmt.Sprintf("%v", schema))
	}
}
