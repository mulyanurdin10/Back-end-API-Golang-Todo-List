package presentation

import "testcode/features/todos"

type TodoHandler struct {
	todoBusiness todos.Business
}

func NewTodoHandler(business todos.Business) *TodoHandler {
	return &TodoHandler{
		todoBusiness: business,
	}
}
