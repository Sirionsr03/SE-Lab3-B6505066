package entity

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	StudentID string
	Year int
	Institute string
	Major string
	PictureStudentID []byte


	MemberID       uint   `gorm:"unique"`
	Products []Products `gorm:"foreignKey:SellerID"`
	
}
