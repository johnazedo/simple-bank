package domain

import(
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:datetime" sql:"index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.NewV4()
	base.CreatedAt = time.Now()
	return nil
}
