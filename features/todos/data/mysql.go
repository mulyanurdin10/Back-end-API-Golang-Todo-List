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

func (repo *mysqlTodoRepository) GetAllData(param string) (data []todos.Core, row int, err error) {
	var getAllData []Todos
	tx := repo.db.Find(&getAllData, param)
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return toCoreList(getAllData), int(tx.RowsAffected), nil
}

func (repo *mysqlTodoRepository) GetData(id int) (data todos.Core, row int, err error) {
	var getData Todos
	tx := repo.db.First(&getData, id)
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return getData.toCore(), int(tx.RowsAffected), nil
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

func (repo *mysqlTodoRepository) UpdateData(id int, insert todos.Core) (data todos.Core, row int, err error) {
	var getData Todos
	tx := repo.db.First(&getData, id).Updates(map[string]interface{}{"title": insert.Title, "is_active": insert.IsActive})
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return getData.toCore(), int(tx.RowsAffected), nil
}

func (repo *mysqlTodoRepository) DeleteData(id int) (row int, err error) {
	var getData Todos
	tx := repo.db.Unscoped().Delete(&getData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
