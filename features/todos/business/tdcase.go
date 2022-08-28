package business

import (
	"errors"
	"testcode/features/todos"

	"github.com/go-playground/validator"
)

type todoTdCase struct {
	todoData todos.Data
}

func NewTodoBusiness(tdData todos.Data) todos.Business {
	return &todoTdCase{
		todoData: tdData,
	}
}

func (tdcase *todoTdCase) InsertData(insert todos.Core) (data todos.Core, row int, err error) {
	v := validator.New()
	errActivitysID := v.Var(insert.ActivitysID, "required")
	if errActivitysID != nil {
		return data, -1, errors.New("activity_group_id cannot be null")
	}
	errTitle := v.Var(insert.Title, "required")
	if errTitle != nil {
		return data, -1, errors.New("title cannot be null")
	}
	insert.Priority = "very-high"
	insert.IsActive = true
	data, row, err = tdcase.todoData.InsertData(insert)
	return data, row, err
}

func (tdcase *todoTdCase) UpdateData(id int, insert todos.Core) (data todos.Core, row int, err error) {
	dataGet, _, _ := tdcase.todoData.GetData(id)
	if insert.Title == "" {
		insert.Title = dataGet.Title
	}
	if insert.IsActive == false && insert.IsActive != true {
		insert.IsActive = dataGet.IsActive
	}
	data, row, err = tdcase.todoData.UpdateData(id, insert)
	return data, row, err
}
