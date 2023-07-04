package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID          int
	Done        bool   `json:"done"`
	Description string `json:"description" validate:"required"`
}

func Create(e echo.Context) error {
	var task Task
	if err := e.Bind(&task); err != nil {
		return e.JSON(http.StatusInternalServerError, "error binding task")
	}
	t, err := readFile()
	if err != nil {
		e.JSON(http.StatusInternalServerError, "error reading file")
	}

	if len(t) == 0 {
		task.ID = 1
	} else {
		index := len(t) - 1
		task.ID = t[index].ID + 1
	}

	t = append(t, task)
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, t)
}

func Get(e echo.Context) error {
	t, err := readFile()
	if err != nil {
		e.JSON(http.StatusInternalServerError, "error reading file")
	}
	return e.JSON(http.StatusOK, t)
}

func Update(e echo.Context) error {
	idP := e.Param("id")
	id, _ := strconv.Atoi(idP)

	fmt.Println(id)
	t, err := readFile()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "error reading file")
	}
	if id > len(t) {
		return e.JSON(http.StatusNotFound, "task not found")
	}

	var i int
	for key, v := range t {
		if v.ID == id {
			i = key
			break
		}
	}

	t[i].Done = true
	data, _ := json.Marshal(t)

	err = ioutil.WriteFile("./tasks.json", data, 0644)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "error store task")
	}

	return e.JSON(http.StatusOK, t)
}

func DeleteAll(e echo.Context) error {
	data := []byte("[]")
	err := ioutil.WriteFile("./tasks.json", data, 0644)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "error store task")
	}

	return e.JSON(http.StatusOK, "All data have been deleted successfully")
}

func readFile() ([]Task, error) {
	file, err := ioutil.ReadFile("./tasks.json")
	var t []Task
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return []Task{}, err

	}

	if len(file) == 0 {
		return []Task{}, errors.New("error reading file")
	}

	err = json.Unmarshal(file, &t)
	if err != nil {
		return []Task{}, errors.New("error reading file")
	}
	return t, nil
}
