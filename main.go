package main

import (
	models "crud_gin/crudproject/connection"
	"crud_gin/crudproject/routes"
	"fmt"
)

func main() {

	db := models.SetupDB()

	r := routes.SetupRoutes(db)
	r.Run(":9020")
	fmt.Println("listening on http://localhost:9020")
}
