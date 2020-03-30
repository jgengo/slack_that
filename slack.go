package main

import (
	"context"
	"errors"
	"log"

	"github.com/slack-go/slack"
	"golang.org/x/time/rate"
)

var s = newSlackTask()

// SlackRequest ...
type SlackRequest struct {
	Workspace   string               `json:"workspace"`
	Channels    []string             `json:"channels"`
	Username    string               `json:"username"`
	AsUser      bool                 `json:"as_user"`
	Parse       string               `json:"parse"`
	LinkNames   int                  `json:"link_names"`
	UnfurlLinks bool                 `json:"unfurl_links"`
	UnfurlMedia bool                 `json:"unfurl_media"`
	IconURL     string               `json:"icon_url"`
	IconEmoji   string               `json:"icon_emoji"`
	Markdown    bool                 `json:"mrkdwn,omitempty"`
	EscapeText  bool                 `json:"escape_text"`
	Text        string               `json:"text"`
	Blocks      []slack.SectionBlock `json:"blocks"`
	Attachments []slack.Attachment   `json:"attachments"`
}

func buildParam(b *SlackRequest) []slack.MsgOption {
	options := []slack.MsgOption{}

	postParameters := slack.PostMessageParameters{
		Username:    b.Username,
		AsUser:      b.AsUser,
		Parse:       b.Parse,
		LinkNames:   b.LinkNames,
		UnfurlLinks: b.UnfurlLinks,
		UnfurlMedia: b.UnfurlMedia,
		IconURL:     b.IconURL,
		IconEmoji:   b.IconEmoji,
		Markdown:    b.Markdown,
		EscapeText:  b.EscapeText,
	}

	options = append(options, slack.MsgOptionPostMessageParameters(postParameters))
	options = append(options, slack.MsgOptionText(b.Text, true))

	if len(b.Blocks) != 0 {
		for _, block := range b.Blocks {
			options = append(options, slack.MsgOptionBlocks(block))
		}
	}

	if len(b.Attachments) != 0 {
		for _, attachment := range b.Attachments {
			options = append(options, slack.MsgOptionAttachments(attachment))
		}
	}

	return options

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

func (s *SlackTask) doSlackTask(channel string, body *SlackRequest, options []slack.MsgOption) {
	s.limit.Wait(context.Background())

	_, _, err := Gateway[body.Workspace].PostMessage(channel, options...)

	if err != nil {
		log.Printf("error while trying to send %v:\n\t->%v", body, err)
	}
}

// ProcessCreate will threat the body and do the job!
func ProcessCreate(body *SlackRequest) error {
	if body.Workspace == "" {
		return errors.New("workspace isn't specified")
	}

	if len(body.Channels) == 0 {
		return errors.New("channel isn't specified")
	}

	if _, ok := Gateway[body.Workspace]; !ok {
		return errors.New("workspace doesn't exist")
	}

	myParam := buildParam(body)

	for _, channel := range body.Channels {
		go s.doSlackTask(channel, body, myParam)
	}

	return nil
}
