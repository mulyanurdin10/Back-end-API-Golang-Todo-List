package request

import "testcode/features/todos"

type Todos struct {
	ActivitysID int    `json:"activity_group_id" form:"activity_group_id"`
	Title       string `json:"title" form:"title"`
	IsActive    bool   `json:"is_active" form:"is_active"`
	Priority    string `json:"priority" form:"priority"`
}

func ToCore(req Todos) todos.Core {
	return todos.Core{
		ActivitysID: req.ActivitysID,
		Title:       req.Title,
		IsActive:    req.IsActive,
		Priority:    req.Priority,
	}
}