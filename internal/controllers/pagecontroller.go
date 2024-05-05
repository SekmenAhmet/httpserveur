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
	if req.Method == http.MethodPost {
		form := p.Serv.ParseFormData(req)
		nom := form["Nom"]
		prenom := form["Prenom"]
		p.Db.Query("INSERT INTO users(name, prenom) VALUES (?, ?)", nom, prenom)
	}
}
