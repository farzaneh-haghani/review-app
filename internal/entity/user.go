package entity

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	
	DisplayName string `json:"display_name"`
	PhoneNumber string `json:"phone_number"`
}

func(u User) Table() string{
	return "users"
}

// type Model struct{
// 	ID        uint `gorm:"primarykey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt DeletedAt `gorm:"index"`
// }