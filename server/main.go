package main

import (
	"fmt"
	"github.com/goldsmithb/spotted_lantern_api/api"
	"github.com/goldsmithb/spotted_lantern_api/config"
	"github.com/goldsmithb/spotted_lantern_api/storage"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
)

func main() {
	fmt.Println("Initializing Lantern Fly API")
	// create zap logger!
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	c, err := config.New("config.yaml", logger)
	if err != nil {
		panic(err)
	}

	db := storage.NewDbClient(c, logger)
	err = db.Connect()
	if err != nil {
		panic(err)
	}

	service := api.NewAPI(c, db)

	server := NewServer(logger, c, service, db)
	server.Start()

}
