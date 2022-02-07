package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var instance *gorm.DB
var locker = &sync.Mutex{}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&ProductModel{})
	if err != nil { panic("Error on ProductModel: " + err.Error()) }
}

func GetInstance() *gorm.DB {
	locker.Lock()
	if instance == nil {
		var err error
		instance, err = gorm.Open(postgres.Open("host=0.0.0.0 port=5432 user=lemonade dbname=products password=lemonade sslmode=disable"), &gorm.Config{})
		if err != nil {
			locker.Unlock()
			panic("Failed to connect database")
		}
	}
	locker.Unlock()
	return instance
}