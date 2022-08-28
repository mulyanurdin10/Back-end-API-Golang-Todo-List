package data

import (
	"testcode/features/activitys"

	"gorm.io/gorm"
)

type Activitys struct {
	gorm.Model
	Email string `gorm:"unique" json:"email" form:"email"`
	Title string `json:"title" form:"title"`
}

func toCoreList(data []Activitys) []activitys.Core {
	result := []activitys.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Activitys) toCore() activitys.Core {
	return activitys.Core{
		ID:        int(data.ID),
		Email:     data.Email,
		Title:     data.Title,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func fromCore(core activitys.Core) Activitys {
	return Activitys{
		Email: core.Email,
		Title: core.Title,
	}
}
