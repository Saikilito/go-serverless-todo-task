package structs

import (
	"github.com/saikilito/go-serverless-todo-task/libs/task/structs"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	firstName string `gorm:"not null"`
	lastName  string `gorm:"not null"`
	email     string `gorm:"not null; unique_index"`
	Task      []structs.Task
}
