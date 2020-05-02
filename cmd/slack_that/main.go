package main

import (
	"log"
	"net/http"

	"github.com/jgengo/slack_that/internal/config"
	"github.com/jgengo/slack_that/internal/router"
)

func main() {
	if err := config.Initiate(); err != nil {
		log.Fatalf("initiate failed: %v\n", err)
	}
	router := router.New()
	http.ListenAndServe("localhost:8080", router)

}
