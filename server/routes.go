package server

import "net/http"

func createRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handleHello)
	return mux
}
