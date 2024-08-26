package main

import (
	"fmt"
	"net/http"
	"simpleWebAppGo/cmd/routes"
	"simpleWebAppGo/cmd/routes/mainRoute"
)

func InitRoutes() {
	mainRoute.MainRender()
}

func main() {
	InitRoutes()

	http.ListenAndServe(":8080", routes.GetMuxInstance())

	fmt.Println("Hello World")
}
