package main

import (
	"fmt"

	"github.com/SQUASHD/boilerplater/internal/server/api"
)

func main() {

	server := api.NewServer()

	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
