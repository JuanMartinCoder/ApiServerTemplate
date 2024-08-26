package mainRoute

import (
	mainDomain "simpleWebAppGo/cmd/domain/main"
	"simpleWebAppGo/cmd/routes"
	"simpleWebAppGo/utils"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc("GET "+utils.RoutesIntance.MAIN, mainDomain.MainView)
}
