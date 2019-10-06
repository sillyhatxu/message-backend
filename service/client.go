package service

import "github.com/sillyhatxu/message-backend/slack"

var birthdaySlackClient *slack.Client

func InitialSlackClient(host string, url string) {
	birthdaySlackClient = slack.NewSlackClient(host, url)
}
