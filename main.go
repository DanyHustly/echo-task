package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Task struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

var tasks = []Task{}

// POST /tasks — создать новую задачу
func createTask(c echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	newTask := Task{
		ID:     uuid.NewString(),
		Task:   req.Task,
		Status: req.Status,
	}

	tasks = append(tasks, newTask)
	return c.JSON(http.StatusCreated, newTask)
}

// GET /tasks — получить все задачи
func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

// PATCH /tasks/:id — обновить задачу по ID
func updateTask(c echo.Context) error {
	id := c.Param("id")
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	for i, t := range tasks {
		if t.ID == id {
			if req.Task != "" {
				tasks[i].Task = req.Task
			}
			if req.Status != "" {
				tasks[i].Status = req.Status
			}
			return c.JSON(http.StatusOK, tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

// DELETE /tasks/:id — удалить задачу по ID
func deleteTask(c echo.Context) error {
	id := c.Param("id")
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Роуты
	e.POST("/tasks", createTask)
	e.GET("/tasks", getTasks)
	e.PATCH("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Start(":8080")
}
