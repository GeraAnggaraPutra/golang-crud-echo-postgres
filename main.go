package main

import (
	"crud-database-postgresql/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start("localhost:5000"))
}