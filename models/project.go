package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Type   string
	Genre  string
	UserID uint
}
