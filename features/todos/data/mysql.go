package data

import (
	"testcode/features/todos"

	"gorm.io/gorm"
)

type mysqlTodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(conn *gorm.DB) todos.Data {
	return &mysqlTodoRepository{
		db: conn,
	}
}
