package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string // this should be a hashed password
	Email    string
}
