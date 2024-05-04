package main

func main() {
	var server HttpServer
	var db Database
	db.Connex()
	var page PageController
	page.db = &db
	server.SetRoute("/", page.Home)
	server.StartServer(8000)
}
