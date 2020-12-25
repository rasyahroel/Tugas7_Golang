package routes

import (
	"echo-rest/controllers"

	"github.com/labstack/echo"
)

// SuppliersRoute ...
func SuppliersRoute(g *echo.Group) {
	g.POST("/list", controllers.FetchAllSuppliers)
	g.POST("/add", controllers.StoreSupplier)
	g.POST("/update", controllers.UpdateSupplier)
	g.POST("/delete", controllers.DeleteSupplier)
}
