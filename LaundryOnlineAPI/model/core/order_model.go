package core

import "gorm.io/gorm"

// Status enum value
const (
	wait_paid = "menunggu-pembayaran"
	paid      = "lunas"
	pick_up   = "jemput"
	process   = "proses"
	deliver   = "diantar"
	done      = "selesai"
)

type Order struct {
	gorm.Model
	CustomerUsername string
	Customer         Customer `gorm:"forignKey:CustomerUsername;not null"`
	ServiceID        uint
	Service          Service `gorm:"forignKey:ServiceID;not null"`
	DryWeight        uint8   `gorm:"not null"`
	TotalPrice       uint32  `gorm:"not null"`
	Status           string  `gorm:"size:50;not null"`
}
