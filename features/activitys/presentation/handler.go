package presentation

import (
	"net/http"
	"testcode/features/activitys"
	"testcode/features/activitys/presentation/request"
	"testcode/features/activitys/presentation/response"

	"github.com/labstack/echo/v4"
)

type ActivityHandler struct {
	activityBusiness activitys.Business
}

func NewActivityHandler(business activitys.Business) *ActivityHandler {
	return &ActivityHandler{
		activityBusiness: business,
	}
}

func (h *ActivityHandler) InsertData(c echo.Context) error {
	var insertData request.Activitys
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to bind data, Check your input",
		})
	}
	newActivity := request.ToCore(insertData)
	data, row, err := h.activityBusiness.InsertData(newActivity)
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
