package routes

import (
	"net/http"

	"github.com/Befous/tutorialstdlibrary/handler"
)

func ContohRoute(app *http.ServeMux) {
	app.HandleFunc("GET /contoh/{z}", handler.ContohHandler)
}
