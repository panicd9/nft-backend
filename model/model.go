package model

import (
	"gorm.io/gorm"
)

type MintedEvent struct {
	gorm.Model
	Minter  string
	TokenId uint64
}

type BlacklistedEvent struct {
	gorm.Model
	AddedBy            string
	BlacklistedAddress string
}

type RegisteredEvent struct {
	gorm.Model
	Address string
}
