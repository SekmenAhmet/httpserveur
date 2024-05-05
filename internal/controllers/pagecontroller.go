package controllers

import (
	"httpserveur/internal/database"
	"httpserveur/internal/httpserver"
	"net/http"
)

type PageController struct {
	Serv *httpserver.HttpServer
	Db   *database.Database
}

func (p *PageController) Home(w http.ResponseWriter, req *http.Request) {
	p.Serv.Render(w, req, "../../internal/views/home.html")
	p.Serv.ServerDetails(w, req)
	form := p.Serv.ParseFormData(req)
	nom := form["Nom"][0]
	prenom := form["Prenom"][0]
	p.Db.Query("Insert into users(name, prenom) values (?, ?)", nom, prenom)
}
