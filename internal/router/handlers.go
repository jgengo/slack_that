package router

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jgengo/slack_that/internal/task"
	"github.com/jgengo/slack_that/internal/utils"
)

// Index is called when it receives a GET on /
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

// Create is called when it receives a new POST on /
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("%shttp (warning)%s while reading the body's request. (%v)\n", utils.Yellow, utils.Reset, err)
		return
	}
	defer r.Body.Close()

	bodyParsed := &task.SlackRequest{}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, bodyParsed); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Error: err.Error()})
		log.Printf("%sjson: (warning)%s while unmarshalling the body's request. (%v)\n", utils.Yellow, utils.Reset, err)
		return
	}

	if err := bodyParsed.ProcessCreate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Error: err.Error()})
		log.Printf("%sjson: (warning)%s while processing ProcessCreate(). (%v)\n", utils.Yellow, utils.Reset, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.SuccessResponse{Success: "request queued"})
}

// Health for health checking the service.
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task.NewHealthResponse())
}
