package model

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	ProfileImage *string  `gorm:"column:name"`
	Description  *string  `gorm:"column:description"`
	Rate         *float64 `gorm:"column:rate"`
	DateOfBirth  *string  `gorm:"column:date_of_birth"`
	Address      *string  `gorm:"column:address"`
	Phone        *string  `gorm:"column:phone"`
}

type User struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Email   string `gorm:"column:email;unique"`
	Role    string `gorm:"column:role"`
	Profile *UserProfile
}
