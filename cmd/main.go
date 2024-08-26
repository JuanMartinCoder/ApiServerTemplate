package main

import (
	"httpServerTemplate/cmd/routes"
	"log"
	"net/http"
)

func main() {
	routes.RenderRoutes()
	err := http.ListenAndServe(":8080", routes.GetMuxInstance())
	if err != nil {
		log.Fatal(err)
	}
}
