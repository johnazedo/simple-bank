package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"news/domain"
)

var db_example[]domain.News;

func DatabaseSetup(){
	db, err := gorm.Open(sqlite.Open("news.db"), &gorm.Config{});

	if err != nil {
		panic("Failed to connect database")
	}

	MigrateData(db, &domain.News{})
}

func MigrateData(db *gorm.DB, schema ...interface{})  {
	err := db.AutoMigrate(schema)
	if err != nil {
		panic("Cannot migrate schema: " + fmt.Sprintf("%v", schema))
	}
}

func GetAllNews() []domain.News {
	db_example = append(db_example, []domain.News{
		{gorm.Model{}, "1", "This is a test", "http://newsletter.com", "johnazedo"},
	}...)

	return db_example
}
