package core

import "gorm.io/gorm"

// Service enum value
const (
	kiloan      = "kiloan"
	satuan      = "satuan"
	VIP         = "VIP"
	karpet      = "karpet"
	setrika_aja = "setrika-aja"
	express     = "express"
)

type Service struct {
	gorm.Model
	Name       string `gorm:"size:11;not null"` // Fill service enum value
	PricePerKg uint16 `gorm:"default:0"`
}
