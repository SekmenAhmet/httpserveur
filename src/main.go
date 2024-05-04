package main

func main() {
	var server HttpServer
	var db Database
	db.Connex() // Initialisation de la connexion à la base de données
	var page PageController
	page.db = &db
	server.SetRoute("/", page.Home)
	server.StartServer(8000)
}
