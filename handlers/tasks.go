package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/reoim/demoApp/models"
	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		result := models.GetTasks(db)
		return c.JSON(http.StatusOK, result)
	}
}



// PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		t := new(models.Task)
		if err := c.Bind(t); err != nil {
			return err
		}
		models.PutTask(db, t.Name)

		return c.JSON(http.StatusCreated, H{
			"created": 123,
		})
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		models.DeleteTask(db, id)
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
