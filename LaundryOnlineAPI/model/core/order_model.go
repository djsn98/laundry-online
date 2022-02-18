package core

import "gorm.io/gorm"

// Status enum value
const (
	wait_paid = "belum-dibayar"
	paid      = "lunas"
	pick_up   = "jemput"
	process   = "proses"
	deliver   = "antar"
	done      = "selesai"
)

type Order struct {
	gorm.Model
	CustomerID uint
	Customer   Customer `gorm:"forignKey:CustomerID"`
	ServiceID  uint
	Service    Service `gorm:"forignKey:ServiceID"`
	DryWeight  uint8   `gorm:"not null"`
	TotalPrice uint32  `gorm:"not null"`
	Status     string  `gorm:"size:13;not null"`
}
