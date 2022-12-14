package data

import (
	"testcode/features/activitys"

	"gorm.io/gorm"
)

type mysqlActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(conn *gorm.DB) activitys.Data {
	return &mysqlActivityRepository{
		db: conn,
	}
}

func (repo *mysqlActivityRepository) GetAllData() (data []activitys.Core, err error) {
	var getAllData []Activities
	tx := repo.db.Find(&getAllData)
	if tx.Error != nil {
		return data, tx.Error
	}
	return toCoreList(getAllData), nil
}

func (repo *mysqlActivityRepository) GetData(id int) (data activitys.Core, row int, err error) {
	var getData Activities
	tx := repo.db.First(&getData, id)
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return getData.toCore(), int(tx.RowsAffected), nil
}

func (repo *mysqlActivityRepository) InsertData(insert activitys.Core) (data activitys.Core, row int, err error) {
	var getData Activities
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData).First(&getData, insertData.ID)
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return getData.toCore(), int(tx.RowsAffected), nil
}

func (repo *mysqlActivityRepository) UpdateData(id int, insert activitys.Core) (data activitys.Core, row int, err error) {
	var getData Activities
	tx := repo.db.First(&getData, id).Updates(map[string]interface{}{"email": insert.Email, "title": insert.Title})
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return getData.toCore(), int(tx.RowsAffected), nil
}

func (repo *mysqlActivityRepository) DeleteData(id int) (row int, err error) {
	var getData Activities
	tx := repo.db.Unscoped().Delete(&getData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlActivityRepository) UniqueData(insert activitys.Core) (row int, err error) {
	var getData Activities
	insertData := fromCore(insert)
	tx := repo.db.Where("email = ?", insertData.Email).First(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
