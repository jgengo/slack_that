package task

import (
	"testing"

	"github.com/jgengo/slack_that/internal/task"

	"github.com/slack-go/slack"
)

func TestCheckIM(t *testing.T) {

	userEmail := "gustavo@42sp.org.br"
	userID := "UMLNGQRFB"
	token := ""

	client := task.SlackClient{
		Value: slack.New(token),
	}

	response, err := client.GetIM(userEmail)

	if err != nil {
		t.Errorf("Error gathering channel ID, got: %s, expected: %s", err, userID)
	}
	if response == "PREVIOUS" {
		t.Errorf("Channel ID Was incorrect, got: %s, expected: %s", response, userID)
	}

}
