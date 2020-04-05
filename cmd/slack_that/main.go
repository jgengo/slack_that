package main

import (
	"net/http"

	"github.com/jgengo/slack_that/internal/config"
	"github.com/jgengo/slack_that/internal/router"
)

func main() {

	config.Initiate()
	router := router.New()
	http.ListenAndServe("localhost:8080", router)

}
