package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// глобальная переменная для хранения задачи
var task string

// структура для разбора JSON из тела запроса
type TaskPayload struct {
	Task string `json:"task"`
}

func main() {
	e := echo.New()

	// маршруты
	e.POST("/task", postTaskHandler)
	e.GET("/", getHelloHandler)

	// запускаем сервер на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}

// POST /task
func postTaskHandler(c echo.Context) error {
	payload := new(TaskPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	task = payload.Task
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

// GET /
func getHelloHandler(c echo.Context) error {
	message := "hello, " + task
	return c.String(http.StatusOK, message)
}
