package main

import (
	"fmt"
	"github.com/goldsmithb/spotted_lantern_api/api"
)

func main() {
	fmt.Println("Hello")

	server := NewServer(nil, api.NewAPI())
	server.Start()

}
