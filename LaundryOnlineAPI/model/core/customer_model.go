package core

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Name      string `gorm:"size:255;not null"`
	Username  string `gorm:"primaryKey;type:varchar(255);not null"` // e.g "@djsn98"
	Order     []Order
	Balance   uint32 `gorm:"default:0"`
	Address   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
