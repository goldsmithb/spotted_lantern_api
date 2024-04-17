package main

import (
	"fmt"
	"github.com/eaigner/jet"
	"github.com/goldsmithb/spotted_lantern_api/api"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello")

	server := NewServer(nil, api.NewAPI())
	go server.Start()

	db, err := jet.Open("postgres", "postgres://ktpsvtav:gNpwtW_yxBDNc3kbTyj3_nBjyPs9fWBh@isilo.db.elephantsql.com/ktpsvtav?sslmode=require")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err == nil {
		fmt.Println(":)")
	}
	var rows []*struct {
		Id       string
		Username string
		Email    string
		Passkey  string
		Score    int
	}

	db.Query("select * from users").Rows(&rows)
}
