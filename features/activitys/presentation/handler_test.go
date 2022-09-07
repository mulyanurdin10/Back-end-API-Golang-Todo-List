package presentation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testcode/features/activitys"
	"testcode/features/activitys/presentation/response"
	"testcode/mocks"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ActivitysResponseSuccess struct {
	Message string
	Data    []response.Activities
}

type ActivitysResponse struct {
	Message string
	Data    response.Activities
}

type ResponseGlobal struct {
	Message string
}

func TestGetAllData(t *testing.T) {
	e := echo.New()
	avtcase := new(mocks.ActivityBusiness)
	GetAllData := []activitys.Core{{ID: 1, Email: "mulyanurdin10@gmail.com", Title: "test-1"}}

	t.Run("Success Get All Data", func(t *testing.T) {
		avtcase.On("GetAllData", mock.Anything).Return(GetAllData, nil).Once()
		srv := NewActivityHandler(avtcase)

		req := httptest.NewRequest(http.MethodGet, "/activity-groups", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/activity-groups")

		responseData := ActivitysResponseSuccess{}
		if assert.NoError(t, srv.GetAllData(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, GetAllData[0].ID, responseData.Data[0].ID)
		}
		avtcase.AssertExpectations(t)
	})

	t.Run("Error Get All Data", func(t *testing.T) {
		avtcase.On("GetAllData", mock.Anything).Return(nil, errors.New("Failed to get all data")).Once()
		srv := NewActivityHandler(avtcase)

		req := httptest.NewRequest(http.MethodGet, "/activity-groups", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/activity-groups")

		srv.GetAllData(echoContext)
		responseBody := rec.Body.String()
		var responseData ResponseGlobal
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Failed to get all data", responseData.Message)
		avtcase.AssertExpectations(t)
	})
}

func TestGetData(t *testing.T) {
	e := echo.New()
	avtcase := new(mocks.ActivityBusiness)
	GetData := activitys.Core{ID: 1, Email: "mulyanurdin10@gmail.com", Title: "test-1"}

	t.Run("Success Get Data", func(t *testing.T) {
		avtcase.On("GetData", mock.Anything).Return(GetData, 1, nil).Once()
		srv := NewActivityHandler(avtcase)

		req := httptest.NewRequest(http.MethodGet, "/activity-groups/:id", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/activity-groups/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ActivitysResponse{}
		if assert.NoError(t, srv.GetData(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, GetData.ID, 1, responseData.Data.ID)
		}
		avtcase.AssertExpectations(t)
	})

	t.Run("Error Get Data Activity with ID Not Found", func(t *testing.T) {
		avtcase.On("GetData", mock.Anything).Return(GetData, 0, errors.New("Activity with ID Not Found")).Once()
		srv := NewActivityHandler(avtcase)

		req := httptest.NewRequest(http.MethodGet, "/activity-groups/:id", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/activity-groups/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("2")

		srv.GetData(echoContext)
		responseBody := rec.Body.String()
		var responseData ResponseGlobal
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Activity with ID 2 Not Found", responseData.Message)
		avtcase.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		avtcase.On("GetData", mock.Anything).Return(GetData, 1, errors.New("Failed to get data")).Once()
		srv := NewActivityHandler(avtcase)

		req := httptest.NewRequest(http.MethodGet, "/activity-groups/:id", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/activity-groups/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		srv.GetData(echoContext)
		responseBody := rec.Body.String()
		var responseData ResponseGlobal
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.NotNil(t, err, "error")
		}
		assert.Equal(t, "Failed to get data", responseData.Message)
		avtcase.AssertExpectations(t)
	})
}
