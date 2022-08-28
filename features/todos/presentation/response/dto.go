package response

import (
	"testcode/features/todos"
	"time"
)

type Todos struct {
	ID              int       `json:"id"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func FromCoreList(data []todos.Core) []Todos {
	result := []Todos{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data todos.Core) Todos {
	return Todos{
		ID:              data.ID,
		ActivityGroupID: data.ActivityGroupID,
		Title:           data.Title,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}
}
