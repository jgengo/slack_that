package main

import (
	"errors"
	"log"
)

// SlackCreate struct to store POST / body
type SlackCreate struct {
	Workspace string `json:"workspace"`  // mandatory
	Channel   string `json:"channel"`    // opt
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

	Gateway["hive-staff"].PostMessage(
		"testing",
		slack.MsgOptionText("test", false),
		slack.MsgOptionIconEmoji(body.IconEmoji),
		slack.MsgOptionUsername(body.Username)
	)

	return nil
}

func displayBody(body *SlackCreate) {
	log.Printf("body: %+v", body)
}
