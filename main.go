package main

import (
	"goback/routes"
	"goback/model"
)

func main() {
	model.InitDb()
	defer model.ExitDb()
	routes.InitRouter()

	
}
