package business

import (
	"errors"
	"testcode/features/activitys"

	"github.com/go-playground/validator"
)

type activityAvtCase struct {
	activityData activitys.Data
}

func NewActivityBusiness(avtData activitys.Data) activitys.Business {
	return &activityAvtCase{
		activityData: avtData,
	}
}

func (avtcase *activityAvtCase) InsertData(insert activitys.Core) (data activitys.Core, row int, err error) {
	v := validator.New()
	errEmail := v.Var(insert.Email, "required,email")
	if errEmail != nil {
		return data, -1, errors.New("invalid format email")
	}
	errTitle := v.Var(insert.Title, "required")
	if errTitle != nil {
		return data, -1, errors.New("title not be null")
	}
	rowUnique, _ := avtcase.activityData.UniqueData(insert)
	if rowUnique == 1 {
		return data, -1, errors.New("email already exists")
	}
	data, row, err = avtcase.activityData.InsertData(insert)
	return data, row, err
}
