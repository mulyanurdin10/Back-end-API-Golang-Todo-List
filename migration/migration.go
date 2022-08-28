package migration

import (
	_mActivitys "testcode/features/activitys/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mActivitys.Activitys{})
}
