package migration

import (
	_mActivitys "testcode/features/activitys/data"
	_mTodos "testcode/features/todos/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mActivitys.Activitys{})
	db.AutoMigrate(&_mTodos.Todos{})
}
