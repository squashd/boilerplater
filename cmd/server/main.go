package main

import (
	"github.com/SQUASHD/boilerplater/internal/server/api"
)

func main() {

	server := api.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
