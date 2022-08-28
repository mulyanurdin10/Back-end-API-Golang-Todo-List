package data

import (
	"testcode/features/todos"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	ActivitysID int    `json:"activity_group_id" form:"activity_group_id"`
	Title       string `json:"title" form:"title"`
	IsActive    bool   `json:"is_active" form:"is_active"`
	Priority    string `json:"priority" form:"priority"`
}

func toCoreList(data []Todos) []todos.Core {
	result := []todos.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Todos) toCore() todos.Core {
	return todos.Core{
		ID:          int(data.ID),
		ActivitysID: data.ActivitysID,
		Title:       data.Title,
		IsActive:    data.IsActive,
		Priority:    data.Priority,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func fromCore(core todos.Core) Todos {
	return Todos{
		ActivitysID: core.ActivitysID,
		Title:       core.Title,
		IsActive:    core.IsActive,
		Priority:    core.Priority,
	}
}
