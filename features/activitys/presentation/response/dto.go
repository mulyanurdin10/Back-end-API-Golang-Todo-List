package response

import (
	"testcode/features/activitys"
	"time"
)

type Activitys struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCoreList(data []activitys.Core) []Activitys {
	result := []Activitys{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data activitys.Core) Activitys {
	return Activitys{
		ID:        data.ID,
		Email:     data.Email,
		Title:     data.Title,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
