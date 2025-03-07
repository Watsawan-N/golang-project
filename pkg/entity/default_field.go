package entity

import (
	"time"

	"gorm.io/gorm"
)

type DefaultField struct {
	gorm.Model
	CreatedAt     time.Time `gorm:"<-:create"`
	CreatedBy     *uint     `gorm:"<-:create"`
	CreatedByUser *User     `gorm:"foreignKey:CreatedBy"`
	UpdatedBy     *uint
	UpdatedByUser *User `gorm:"foreignKey:UpdatedBy"`
}
