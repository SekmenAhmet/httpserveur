package main

func main() {
	var server HttpServer
	var page PageController
	server.SetRoute("/", page.Home)
	server.StartServer(8000)
}
