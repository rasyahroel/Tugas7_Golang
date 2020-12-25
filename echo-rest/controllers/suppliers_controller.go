package controllers

import (
	"echo-rest/models"
	"net/http"

	"github.com/labstack/echo"
)

// FetchAllSuppliers ... Get All Suppliers
func FetchAllSuppliers(c echo.Context) (err error) {
	result, err := models.FetchSuppliers()
	return c.JSON(http.StatusOK, result)
}

// StoreSupplier ... Insert Supplier
func StoreSupplier(c echo.Context) (err error) {
	result, err := models.StoreSupplier(c)
	return c.JSON(http.StatusOK, result)
}

// UpdateSupplier ...
func UpdateSupplier(c echo.Context) (err error) {
	result, err := models.UpdateSupplier(c)
	return c.JSON(http.StatusOK, result)
}

// DeleteSupplier ...
func DeleteSupplier(c echo.Context) (err error) {
	result, err := models.DeleteSupplier(c)
	return c.JSON(http.StatusOK, result)
}
