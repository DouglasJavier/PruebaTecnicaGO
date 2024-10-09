package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(100);not null"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}
