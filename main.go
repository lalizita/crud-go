package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/lalizita/crud-go/routes"
)

func main() {
	e := echo.New()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8001"))
	fmt.Println("hello world")
}
