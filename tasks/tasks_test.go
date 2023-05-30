package tasks_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lalizita/crud-go/tasks"
)

func TestCreate(t *testing.T) {
	t.Log("Testing create function")
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/create", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := tasks.Create(c)
	if err != nil {
		t.Error("Error calling create function")
	}
}
