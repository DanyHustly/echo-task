package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// 1. Глобальная переменная task
var task string

// 2. Структура запроса для POST
type TaskRequest struct {
	Task string `json:"task"`
}

// 3. Обработчик POST — получает JSON и сохраняет task
func postTask(c echo.Context) error {
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
	}

	// Сохраняем значение task
	task = req.Task

	return c.JSON(http.StatusOK, map[string]string{"message": "task saved"})
}

// 4. Обработчик GET — возвращает приветствие с task
func getTask(c echo.Context) error {
	return c.String(http.StatusOK, "hello, "+task)
}

func main() {
	e := echo.New()

	// Middleware (по желанию)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Роуты
	e.POST("/task", postTask)
	e.GET("/task", getTask)

	// Запуск сервера
	e.Start(":8080")
}
