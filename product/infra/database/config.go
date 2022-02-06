package database

import (
	"gorm.io/driver/sqlite"
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
		instance, err = gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
		if err != nil {
			locker.Unlock()
			panic("Failed to connect database")
		}
	}
	locker.Unlock()
	return instance
}