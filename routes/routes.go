package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lalizita/crud-go/tasks"
)

func Init(e *echo.Echo) {
	e.POST("/create", tasks.Create)
}
