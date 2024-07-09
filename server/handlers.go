package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Msg string `json:"msg"`
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	p := Payload{"Hello World Air"}
	w.Header().Set("Content-Type", "application.json")
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error detected: %v", err)
	}
	w.Write(j)
}
