package database

import (
	. "github.com/JohnAzedo/eCommerce/product/domain"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID        string `gorm:"primarykey;autoIncrement:false;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ProductModel struct {
	Model
	Name        string
	Price       float32
	Description string
	Measurement Measurement
	Amount      int
}

func (pm ProductModel) ToProductEntity() *Product{
	return &Product{
		UUID: pm.UUID,
		Name: pm.Name,
		Price: pm.Price,
		Description: pm.Description,
		Measurement: pm.Measurement,
		Amount: pm.Amount,
	}
}

func FromProductEntity(entity *Product) *ProductModel {
	return &ProductModel{
		Model: Model{ UUID: entity.UUID },
		Name: entity.Name,
		Price: entity.Price,
		Description: entity.Description,
		Measurement: entity.Measurement,
		Amount: entity.Amount,
	}
}