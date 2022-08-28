package presentation

import "testcode/features/activitys"

type ActivityHandler struct {
	activityBusiness activitys.Business
}

func NewActivityHandler(business activitys.Business) *ActivityHandler {
	return &ActivityHandler{
		activityBusiness: business,
	}
}
