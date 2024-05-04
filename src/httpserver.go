package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type HttpServer struct {
	Port int
	Post map[string][]string
	Get  []byte
}

func (h *HttpServer) StartServer(port int) {
	h.Port = port
	fmt.Printf("Serveur démarrer sur le port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (h *HttpServer) ServerDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s<br>", displayData(h.parseFormData(r)))
	fmt.Fprintf(w, "Méthode : %v<br>", h.GetMethod(r))
	fmt.Fprintf(w, "URL : %v<br>", h.GetURL(r))
	fmt.Fprintf(w, "PORT : %v<br>", h.GetPort(r))
	fmt.Fprintf(w, "HOST : %v<br>", h.GetHost(r))
	fmt.Fprintf(w, "Protocol : %v<br>", h.GetProtocol(r))
	h.Logs(r)
}

func displayData(postData map[string][]string) string {
	result := "Form : <br>"
	for key, values := range postData {
		result += fmt.Sprintf("%s : ", key)
		for _, value := range values {
			result += fmt.Sprintf("%s<br>", value)
		}
	}
	return result
}

func (h *HttpServer) parseFormData(r *http.Request) map[string][]string {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		return nil
	}
	if err := r.ParseForm(); err != nil {
		log.Println("Erreur lors de l'analyse du formulaire:", err)
		return nil
	}
	return r.Form
}

func (h *HttpServer) GetMethod(r *http.Request) string {
	return r.Method
}

func (h *HttpServer) GetURL(r *http.Request) string {
	return r.URL.Path
}

func (h *HttpServer) GetPort(r *http.Request) int {
	return h.Port
}

func (h *HttpServer) GetHost(r *http.Request) string {
	return r.Host
}
func (h *HttpServer) GetProtocol(r *http.Request) string {
	return r.Proto
}

func (h *HttpServer) Logs(r *http.Request) {
	log.Printf("Requête %v sur l'URL %v", r.Method, r.URL.Path)
}

func (h *HttpServer) Render(w http.ResponseWriter, r *http.Request, filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur de lecture du fichier: %s", err), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur d'écriture dans la réponse: %s", err), http.StatusInternalServerError)
		return
	}
}

func (h *HttpServer) SetRoute(url string, handler http.HandlerFunc) {
	http.HandleFunc(url, handler)
}
