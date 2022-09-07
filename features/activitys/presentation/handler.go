package presentation

import (
	"net/http"
	"strconv"
	"testcode/features/activitys"
	"testcode/features/activitys/presentation/request"
	"testcode/features/activitys/presentation/response"
	"testcode/helper"

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

func (h *ActivityHandler) GetAllData(c echo.Context) error {
	data, err := h.activityBusiness.GetAllData()
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

func (h *ActivityHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idActivity, errId := strconv.Atoi(id)
	strIdActivity := strconv.Itoa(idActivity)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Invalid ID",
		})
	}
	data, row, err := h.activityBusiness.GetData(idActivity)
	if row == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "Not Found",
			"message": "Activity with ID " + strIdActivity + " Not Found",
			"data":    map[string]interface{}{},
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success", data))
}

func (h *ActivityHandler) InsertData(c echo.Context) error {
	var insertData request.Activities
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

func (h *ActivityHandler) UpdateData(c echo.Context) error {
	id := c.Param("id")
	idActivity, errId := strconv.Atoi(id)
	strIdActivity := strconv.Itoa(idActivity)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Invalid ID",
		})
	}
	email := c.FormValue("email")
	title := c.FormValue("title")
	var insertData = request.Activities{
		Email: email,
		Title: title,
	}
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to bind data, Check your input",
		})
	}
	newActivity := request.ToCore(insertData)
	data, row, err := h.activityBusiness.UpdateData(idActivity, newActivity)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": err.Error(),
		})
	}
	if row == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "Not Found",
			"message": "Activity with ID " + strIdActivity + " Not Found",
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

func (h *ActivityHandler) DeleteData(c echo.Context) error {
	id := c.Param("id")
	idActivity, errId := strconv.Atoi(id)
	strIdActivity := strconv.Itoa(idActivity)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Invalid ID",
		})
	}
	row, err := h.activityBusiness.DeleteData(idActivity)
	if row == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "Not Found",
			"message": "Activity with ID " + strIdActivity + " Not Found",
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
