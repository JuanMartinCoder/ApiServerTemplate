package mainDomain

import (
	"fmt"
	"net/http"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}