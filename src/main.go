package main

func main() {
	var server HttpServer
	var page PageController
	var db Database
	db.Connex()
	server.SetRoute("/", page.Home)
	server.StartServer(8000)
}
