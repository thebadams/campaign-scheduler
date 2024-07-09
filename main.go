package main

import (
	"fmt"

	"github.com/thebadams/campaign-scheduler/server"
)

func main() {
	s := server.CreateServer()
	fmt.Println("Server Starting...")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Server Could Not Start: %v", err)
	}
}
