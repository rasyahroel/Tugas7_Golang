package controllers

import (
	"echo-rest/models"
	"net/http"

	"github.com/labstack/echo"
)

// FetchAllEmployees ... Get All Employees
func FetchAllEmployees(c echo.Context) (err error) {
	result, err := models.FetchEmployees()
	return c.JSON(http.StatusOK, result)
}

// StoreEmployee ... Insert Employee
func StoreEmployee(c echo.Context) (err error) {
	result, err := models.StoreEmployee(c)
	return c.JSON(http.StatusOK, result)
}

// UpdateEmployee ...
func UpdateEmployee(c echo.Context) (err error) {
	result, err := models.UpdateEmployee(c)
	return c.JSON(http.StatusOK, result)
}

// DeleteEmployee ...
func DeleteEmployee(c echo.Context) (err error) {
	result, err := models.DeleteEmployee(c)
	return c.JSON(http.StatusOK, result)
}
