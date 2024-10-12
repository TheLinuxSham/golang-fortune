package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TheLinuxSham/golang-fortune/internal/fortune"
)

func Root(w http.ResponseWriter, r *http.Request) {
	type responseJson struct {
		Animal string `json:"animal"`
		Quote  string `json:"quote"`
	}

	a, q := fortune.FortuneTeller()

	res := responseJson{
		Animal: a,
		Quote:  q,
	}

	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error rendering in root handler: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
