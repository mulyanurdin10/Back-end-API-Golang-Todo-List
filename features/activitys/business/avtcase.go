package business

import "testcode/features/activitys"

type activityAvtCase struct {
	activityData activitys.Data
}

func NewActivityBusiness(avtData activitys.Data) activitys.Business {
	return &activityAvtCase{
		activityData: avtData,
	}
}
