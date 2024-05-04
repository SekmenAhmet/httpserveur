package main

import "net/http"

type PageController struct {
}

func (p *PageController) Home(w http.ResponseWriter, req *http.Request) {
	var server HttpServer
	server.Render(w, req, "../pages/home.html")
	server.ServerDetails(w, req)
}

func (p *PageController) Test(w http.ResponseWriter, req *http.Request) {
	var server HttpServer
	server.Render(w, req, "../pages/test.html")
}

func (p *PageController) Profil(w http.ResponseWriter, req *http.Request) {
	var server HttpServer
	server.Render(w, req, "../pages/profil.html")
}
