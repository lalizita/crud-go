package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lalizita/crud-go/tasks"
)

func Init(e *echo.Echo) {
	e.GET("/tasks", tasks.Get)
	e.POST("/tasks/create", tasks.Create)
	e.PUT("/tasks/done/:id", tasks.Update)
	e.DELETE("/tasks/delete/all", tasks.DeleteAll)
}
