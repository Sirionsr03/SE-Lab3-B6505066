package entity

import (

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username string
	Password string
	Email string
	FirstName string
	LastName string
	PhoneNumber string
	Address string
	PicProfile []byte

	// SellerID uint
	Seller  Seller `gorm:"foreignKey:MemberID"`

}
