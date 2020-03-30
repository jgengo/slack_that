package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// Index is called when it receives a GET on /
func Index(w http.ResponseWriter, r *http.Request) {
	html := ""
	e := reflect.ValueOf(&SlackRequest{}).Elem()
	fmt.Println(e)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type

		html += fmt.Sprintf("<tr><td style='width: 12em;'>%v</td><td>%v</td></tr>\n", varName, varType)
	}

	fmt.Fprintln(w, "Welcome!<br /><br /><table>"+html+"</table>")
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Panicf("http/json: (error) while Unmarshall body: %v\n", err) // TODO: review this
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panicf("http/json: (error) while encoding the error: %v\n", err)
		}
	}

	if err := ProcessCreate(&bodyParsed); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Panicf("http/json: (error) while processing ProcessCreate: %v\n", err) // TODO: review this
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panicf("http/json: (error) while encoding the error: %v\n", err)
		}
	}
}
