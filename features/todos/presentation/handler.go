package presentation

import (
	"net/http"
	"testcode/features/todos"
	"testcode/features/todos/presentation/request"
	"testcode/features/todos/presentation/response"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoBusiness todos.Business
}

func NewTodoHandler(business todos.Business) *TodoHandler {
	return &TodoHandler{
		todoBusiness: business,
	}
}

func (h *TodoHandler) InsertData(c echo.Context) error {
	var insertData request.Todos
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to bind data, Check your input",
		})
	}
	newTodo := request.ToCore(insertData)
	data, row, err := h.todoBusiness.InsertData(newTodo)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Bad Request",
			"message": err.Error(),
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Data failed to save",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "Success",
		"message": "Success",
		"data":    response.FromCore(data),
	})
}
