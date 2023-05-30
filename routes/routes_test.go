package routes_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lalizita/crud-go/routes"
)

var createJSON = `{"task":"wash the dishes","done":false}`

func TestCreateTask(t *testing.T) {
	e := echo.New()
	routes.Init(e)
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(createJSON))
	rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	e.Router()

	t.Log("=============")
	t.Log(x)
	// if assert.NoError(t,)
}
