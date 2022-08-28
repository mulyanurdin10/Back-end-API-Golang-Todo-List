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
