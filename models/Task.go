package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task         string          `gorm:"type:varchar(250);not null"`
	UserID       uint            `gorm:"not null"`
	TaskCategory []Task_Category `gorm:"foreignKey:TaskID"`
}
