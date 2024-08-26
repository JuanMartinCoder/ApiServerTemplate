package mainRoute

import (
	controller "simpleWebAppGo/cmd/controllers"
	"simpleWebAppGo/cmd/routes"
	"simpleWebAppGo/utils"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc("GET "+utils.RoutesIntance.MAIN, controller.MainView)
}
