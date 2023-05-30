package tasks

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Create(e echo.Context) error {
	fmt.Println("CREATE")
	return nil
}
