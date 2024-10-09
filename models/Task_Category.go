package models

import "gorm.io/gorm"

type Task_Category struct {
	gorm.Model
	TaskID     uint
	CategoryID uint
}
