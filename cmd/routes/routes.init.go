package routes

import (
	C "httpServerTemplate/cmd/controllers"
	"httpServerTemplate/utils"
	"net/http"
)

var mux = http.NewServeMux()

func GetMuxInstance() *http.ServeMux {
	return mux
}

func RenderRoutes() {
	GetMuxInstance().HandleFunc("GET "+utils.RoutesIntance.MAIN, C.MainView)
	GetMuxInstance().HandleFunc("GET "+utils.RoutesIntance.ERROR, C.ErrorView)
}
