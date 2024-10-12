package server

import (
	"net/http"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Root)

	server := &http.Server{
		Addr:    ":3030",
		Handler: mux,
	}

	return server
}
