package controllers

import (
	"encoding/json"
	"fmt"
	"httpServerTemplate/utils"
	"io"
	"net/http"
)

var api_url string = "https://dolarapi.com/v1/dolares"

func getData() ([]utils.Dolar, error) {
	res, err := http.Get(api_url)
	if err != nil {
		return []utils.Dolar{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []utils.Dolar{}, err
	}

	var Dolares []utils.Dolar
	if err := json.Unmarshal(data, &Dolares); err != nil {
		return []utils.Dolar{}, err
	}

	return Dolares, nil
}

func MainView(w http.ResponseWriter, r *http.Request) {
	Dolares, err := getData()
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Dolares)
}
