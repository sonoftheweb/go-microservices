package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email;unique"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role"`
}
