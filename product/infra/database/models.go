package database

import (
	. "github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID        uuid.UUID `gorm:"primarykey;autoIncrement:false"`
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