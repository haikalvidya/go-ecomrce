package seeders

import (
	"go-ecomrce/models"

	"github.com/jinzhu/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
}

func (userSeeder *UserSeeder) SetUsers() {
	users := map[int]map[string]string{
		1: {
			"email":    "user1@things.com",
			"name":     "user1",
			"password": "password1234",
		},
		2: {
			"email":    "user2@things.com",
			"name":     "user2",
			"password": "password212312",
		},
	}

	for key, value := range users {
		user := models.User{}
		userSeeder.DB.First(&user, key)

		if user.ID == 0 {
			user.ID = uint(key)
			user.Email = value["email"]
			user.Name = value["name"]
			user.Password = value["password"]
			userSeeder.DB.Create(&user)
		}
	}
}
