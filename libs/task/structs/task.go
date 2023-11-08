package structs

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	title       string `gorm:"not null;unique_index"`
	description string
	done        bool `gorm:"default:false"`
	userId      uint
}
