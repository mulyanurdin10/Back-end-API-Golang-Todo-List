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

func (repo *mysqlTodoRepository) InsertData(insert todos.Core) (data todos.Core, row int, err error) {
	var getData Todos
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData).First(&getData, insertData.ID)
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return getData.toCore(), int(tx.RowsAffected), nil
}
