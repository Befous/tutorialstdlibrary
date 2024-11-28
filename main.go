package main

import (
	"fmt"
	"net/http"

	"github.com/Befous/tutorialstdlibrary/middleware"
	"github.com/Befous/tutorialstdlibrary/routes"
)

func main() {
	app := http.NewServeMux()
	routes.ContohRoute(app)
	server := http.Server{
		Addr:    ":3000",
		Handler: middleware.AllowCors(app),
	}
	fmt.Println("Starting server on port :3000")
	server.ListenAndServe()
}
