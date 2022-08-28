package presentation

import (
	"net/http"
	"strconv"
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

func (h *TodoHandler) GetAllData(c echo.Context) error {
	match := c.QueryParam("activity_group_id")
	data, row, err := h.todoBusiness.GetAllData(match)
	if row == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    response.FromCoreList(data),
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Success",
		"data":    response.FromCoreList(data),
	})
}

func (h *TodoHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idTodo, errId := strconv.Atoi(id)
	strIdTodo := strconv.Itoa(idTodo)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Invalid ID",
		})
	}
	data, row, err := h.todoBusiness.GetData(idTodo)
	if row == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "Not Found",
			"message": "Todo with ID " + strIdTodo + " Not Found",
			"data":    map[string]interface{}{},
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Success",
		"data":    response.FromCore(data),
	})
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

func (h *TodoHandler) UpdateData(c echo.Context) error {
	id := c.Param("id")
	idTodo, errId := strconv.Atoi(id)
	strIdTodo := strconv.Itoa(idTodo)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Invalid ID",
		})
	}
	title := c.FormValue("title")
	boolIsActive, _ := strconv.ParseBool(c.FormValue("is_active"))
	var insertData = request.Todos{
		Title:    title,
		IsActive: boolIsActive,
	}
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to bind data, Check your input",
		})
	}
	newTodo := request.ToCore(insertData)
	data, row, err := h.todoBusiness.UpdateData(idTodo, newTodo)
	if row == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "Not Found",
			"message": "Todo with ID " + strIdTodo + " Not Found",
			"data":    map[string]interface{}{},
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "data failed to change",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Success",
		"data":    response.FromCore(data),
	})
}

func (h *TodoHandler) DeleteData(c echo.Context) error {
	id := c.Param("id")
	idTodo, errId := strconv.Atoi(id)
	strIdTodo := strconv.Itoa(idTodo)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Invalid ID",
		})
	}
	row, err := h.todoBusiness.DeleteData(idTodo)
	if row == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "Not Found",
			"message": "Todo with ID " + strIdTodo + " Not Found",
			"data":    map[string]interface{}{},
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Success",
		"data":    map[string]interface{}{},
	})
}
