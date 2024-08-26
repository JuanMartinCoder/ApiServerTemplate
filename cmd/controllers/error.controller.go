package controllers

import (
	"fmt"
	"net/http"
)

func ErrorView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ErrorPage")
}
