package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	server := NewServer(nil, nil)
	server.Start()

}
