package database

import (
	. "github.com/JohnAzedo/eCommerce/product/infra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormRepository interface {
	OpenDatabase() (db *gorm.DB, err error)
}

type GormRepositoryImpl struct {
	GormRepository
}

func (repo GormRepositoryImpl) OpenDatabase() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	return db, err
}

func (repo GormRepositoryImpl) Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&ProductModel{})
	if err != nil { panic("Error on ProductModel: " + err.Error()) }
}
