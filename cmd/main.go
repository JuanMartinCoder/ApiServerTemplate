package main

import (
	"fmt"
	"httpServerTemplate/cmd/routes"
	"net/http"
)

func main() {
	routes.RenderRoutes()

	http.ListenAndServe(":8080", routes.GetMuxInstance())

	fmt.Println("Hello World")
}
