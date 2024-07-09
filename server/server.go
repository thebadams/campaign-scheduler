package server

import "net/http"

func CreateServer() *http.Server {
	router := createRoutes()
	s := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	return s
}
