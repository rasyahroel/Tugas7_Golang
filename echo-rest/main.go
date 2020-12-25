package main

import (
	db "echo-rest/db"
	routes "echo-rest/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":3000"))
}
