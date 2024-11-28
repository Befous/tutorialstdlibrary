package routes

import (
	"net/http"

	"github.com/Befous/tutorialstdlibrary/handler"
)

func ContohRoute(app *http.ServeMux) {
	app.HandleFunc("GET /routeget/{z}", handler.ContohHandler)
	app.HandleFunc("POST /routepost/{z}", handler.ContohHandler)
	app.HandleFunc("PUT /routeput/{z}", handler.ContohHandler)
	app.HandleFunc("DELETE /routedelete/{z}", handler.ContohHandler)
}
