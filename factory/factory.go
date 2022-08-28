package factory

import (
	_activityBusiness "testcode/features/activitys/business"
	_activityData "testcode/features/activitys/data"
	_activityPresentation "testcode/features/activitys/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	ActivityPresenter *_activityPresentation.ActivityHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	activityData := _activityData.NewActivityRepository(dbConn)
	activityBusiness := _activityBusiness.NewActivityBusiness(activityData)
	activityPresentation := _activityPresentation.NewActivityHandler(activityBusiness)

	return Presenter{
		ActivityPresenter: activityPresentation,
	}
}
