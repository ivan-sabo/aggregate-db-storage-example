package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string
	AddressCollection []Address
}

type Address struct {
	gorm.Model
	City   string
	UserID uint
}
