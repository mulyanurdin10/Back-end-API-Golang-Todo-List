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

func (repo *mysqlActivityRepository) InsertData(insert activitys.Core) (data activitys.Core, row int, err error) {
	var getData Activitys
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData).First(&getData, insertData.ID)
	if tx.Error != nil {
		return data, 0, tx.Error
	}
	return getData.toCore(), int(tx.RowsAffected), nil
}

func (repo *mysqlActivityRepository) UniqueData(insert activitys.Core) (row int, err error) {
	var getData Activitys
	insertData := fromCore(insert)
	tx := repo.db.Where("email = ?", insertData.Email).First(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
