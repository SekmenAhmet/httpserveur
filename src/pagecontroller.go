package main

import (
	"net/http"
)

type PageController struct {
	Serv *HttpServer
	db   *Database
}

func (p *PageController) Home(w http.ResponseWriter, req *http.Request) {
	p.Serv.Render(w, req, "../pages/home.html")
	formData := p.Serv.parseFormData(req)
	nom := formData["Nom"]
	prenom := formData["Prenom"]
	err := p.db.Query("INSERT INTO users (name, prenom) VALUES (?, ?)", nom[0], prenom[0])
	if err != nil {
		http.Error(w, "Erreur lors de l'insertion dans la base de données: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PageController) Test(w http.ResponseWriter, req *http.Request) {
	p.Serv.Render(w, req, "../pages/test.html")
}

func (p *PageController) Profil(w http.ResponseWriter, req *http.Request) {
	p.Serv.Render(w, req, "../pages/profil.html")
}
