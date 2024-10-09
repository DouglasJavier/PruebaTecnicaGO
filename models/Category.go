package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category     string          `gorm:"type:varchar(250);not null"`
	TaskCategory []Task_Category `gorm:"foreignKey:CategoryID"`
}
