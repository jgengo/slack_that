package main

import (
	"context"
	"errors"
	"log"

	"github.com/slack-go/slack"
	"golang.org/x/time/rate"
)

var s = newSlackTask()

// SlackCreate struct to store POST / body
type SlackCreate struct {
	Workspace string `json:"workspace"`  // mandatory
	Channel   string `json:"channel"`    // mandatory
	Text      string `json:"text"`       // ...
	IconEmoji string `json:"icon_emoji"` // opt
	Username  string `json:"username"`   // opt
}

// SlackTask ...
type SlackTask struct {
	limit *rate.Limiter
}

func newSlackTask() *SlackTask {
	return &SlackTask{
		limit: rate.NewLimiter(1, 1),
	}
}

func (s *SlackTask) doSlackTask(body *SlackCreate) {
	s.limit.Wait(context.Background())
	resp, _, err := Gateway[body.Workspace].PostMessage(
		body.Channel,
		slack.MsgOptionText(body.Text, false),
		slack.MsgOptionIconEmoji(body.IconEmoji),
		slack.MsgOptionUsername(body.Username),
	)
	if err != nil {
		log.Printf("error while trying to send %v:\n\t->%v", body, err)
	}
	log.Println(resp)
}

// ProcessCreate will threat the body and do the job!
func ProcessCreate(body *SlackCreate) error {
	displayBody(body)

	if body.Workspace == "" {
		return errors.New("workspace isn't specified")
	}

	if body.Channel == "" {
		return errors.New("channel isn't specified")
	}

	if _, ok := Gateway[body.Workspace]; !ok {
		return errors.New("workspace doesn't exist")
	}

	go s.doSlackTask(body)

	return nil
}

func displayBody(body *SlackCreate) {
	log.Printf("body: %+v", body)
}
