package factory

import (
	_activityBusiness "testcode/features/activitys/business"
	_activityData "testcode/features/activitys/data"
	_activityPresentation "testcode/features/activitys/presentation"

	_todoBusiness "testcode/features/todos/business"
	_todoData "testcode/features/todos/data"
	_todoPresentation "testcode/features/todos/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	ActivityPresenter *_activityPresentation.ActivityHandler
	TodoPresenter     *_todoPresentation.TodoHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	activityData := _activityData.NewActivityRepository(dbConn)
	activityBusiness := _activityBusiness.NewActivityBusiness(activityData)
	activityPresentation := _activityPresentation.NewActivityHandler(activityBusiness)

	todoData := _todoData.NewTodoRepository(dbConn)
	todoBusiness := _todoBusiness.NewTodoBusiness(todoData)
	todoPresentation := _todoPresentation.NewTodoHandler(todoBusiness)

	return Presenter{
		ActivityPresenter: activityPresentation,
		TodoPresenter:     todoPresentation,
	}
}
