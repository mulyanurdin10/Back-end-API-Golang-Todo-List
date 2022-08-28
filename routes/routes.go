package routes

import (
	"testcode/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))

	e.Pre(middleware.RemoveTrailingSlash())

	// Activitys
	e.GET("/activity-groups", presenter.ActivityPresenter.GetAllData)
	e.GET("/activity-groups/:id", presenter.ActivityPresenter.GetData)
	e.POST("/activity-groups", presenter.ActivityPresenter.InsertData)
	e.PATCH("/activity-groups/:id", presenter.ActivityPresenter.UpdateData)
	e.DELETE("/activity-groups/:id", presenter.ActivityPresenter.DeleteData)

	// Todos
	e.GET("/todo-items", presenter.TodoPresenter.GetAllData)
	e.GET("/todo-items/:id", presenter.TodoPresenter.GetData)
	e.POST("/todo-items", presenter.TodoPresenter.InsertData)
	e.PATCH("/todo-items/:id", presenter.TodoPresenter.UpdateData)
	e.DELETE("/todo-items/:id", presenter.TodoPresenter.DeleteData)

	return e
}
