package main

import (
	"httpserveur/internal/controllers"
	"httpserveur/internal/database"
	"httpserveur/internal/httpserver"
)

func main() {
	db := database.NewDatabase()
	db.Connex()

	serv := &httpserver.HttpServer{}
	page := controllers.PageController{
		Serv: serv,
		Db:   db,
	}

	serv.SetRoute("/", page.Home)

	serv.StartServer(8000)
}
