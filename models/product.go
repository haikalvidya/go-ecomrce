package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:text"`
	Description string `json:"desc" gorm:"type:text"`
	Price       uint   `json:"price" gorm:"type:uint"`
	ImageUrl    string `json:"image_url" gorm:"type:text"`
	UserID      uint
	User        User `gorm:"foreignkey:UserID"`
}
