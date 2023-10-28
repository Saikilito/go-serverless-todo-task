package database

import (
	taskStructs "github.com/saikilito/go-serverless-todo-task/libs/task/structs"
	userStructs "github.com/saikilito/go-serverless-todo-task/libs/user/structs"
)

func dbAutomigrate() {

	DB.AutoMigrate(taskStructs.Task{})
	DB.AutoMigrate(userStructs.User{})

}
