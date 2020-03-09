package main

import (
	"errors"
	"log"

	"github.com/slack-go/slack"
)

// SlackCreate struct to store POST / body
type SlackCreate struct {
	Workspace string `json:"workspace"`  // mandatory
	Channel   string `json:"channel"`    // mandatory
	Text      string `json:"text"`       // ...
	IconEmoji string `json:"icon_emoji"` // opt
	Username  string `json:"username"`   // opt
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

	Gateway[body.Workspace].PostMessage(
		body.Channel,
		slack.MsgOptionText(body.Text, false),
		slack.MsgOptionIconEmoji(body.IconEmoji),
		slack.MsgOptionUsername(body.Username),
	)

	return nil
}

func displayBody(body *SlackCreate) {
	log.Printf("body: %+v", body)
}
