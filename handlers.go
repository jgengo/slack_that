package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Index is called when it receives a GET on /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// Create is called when it receives a new POST on /
func Create(w http.ResponseWriter, r *http.Request) {
	var bodyParsed SlackRequest

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("net/http (error) while reading body: %v\n", err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panicf("net/http: (error) while closing the reader: %v\n", err)
	}

	if err := json.Unmarshal(body, &bodyParsed); err != nil {
		log.Panicf("http/json: (error) while Unmarshall body: %v\n", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panicf("http/json: (error) while encoding the error: %v\n", err)
		}
	}

	if err := ProcessCreate(&bodyParsed); err != nil {
		log.Panicf("http/json: (error) while processing ProcessCreate: %v\n", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panicf("http/json: (error) while encoding the error: %v\n", err)
		}
	}
}
