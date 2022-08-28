package business

import "testcode/features/todos"

type todoTdCase struct {
	todoData todos.Data
}

func NewTodoBusiness(tdData todos.Data) todos.Business {
	return &todoTdCase{
		todoData: tdData,
	}
}
