package routes

import (
	"echo-rest/controllers"

	"github.com/labstack/echo"
)

//UsersRoute is ...
func UsersRoute(g *echo.Group) {
	g.POST("/list", controllers.FetchAllUsers)
	g.POST("/add", controllers.StoreUser)
	g.POST("/update", controllers.UpdateUser)
	g.POST("/delete", controllers.DeleteUser)
	g.POST("/login", controllers.CheckLogin)
}
