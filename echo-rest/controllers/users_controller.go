package controllers

import (
	"echo-rest/models"
	"net/http"

	"github.com/labstack/echo"
)

// FetchAllUsers ... Get All Users
func FetchAllUsers(c echo.Context) (err error) {
	result, err := models.FetchUsers()
	return c.JSON(http.StatusOK, result)
}

// StoreUser ...
func StoreUser(c echo.Context) (err error) {
	result, err := models.StoreUser(c)
	return c.JSON(http.StatusOK, result)
}

// UpdateUser ...
func UpdateUser(c echo.Context) (err error) {
	result, err := models.UpdateUser(c)
	return c.JSON(http.StatusOK, result)
}

// DeleteUser ...
func DeleteUser(c echo.Context) (err error) {
	result, err := models.DeleteUser(c)
	return c.JSON(http.StatusOK, result)
}
