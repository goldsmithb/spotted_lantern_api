package main

import (
	"fmt"
	"github.com/goldsmithb/spotted_lantern_api/api"
	"github.com/goldsmithb/spotted_lantern_api/config"
	"github.com/goldsmithb/spotted_lantern_api/storage"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello")

	c, err := config.New("config.yaml", nil)
	if err != nil {
		panic(err)
	}

	db := storage.NewDbClient(c)
	err = db.Connect()
	if err != nil {
		panic(err)
	}

	service := api.NewAPI(c, db)

	server := NewServer(nil, c, service)
	server.Start()

}
