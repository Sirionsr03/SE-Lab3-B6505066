package entity

import (

	"gorm.io/gorm"
)

type Products struct {
		gorm.Model
		Title string
		Description string
		Price float32
		Picproduct []byte
		Condition string
		Weigth float32
		Status string

		SellerID uint
		Seller   Seller `gorm:"foreignKey:SellerID"`

}