package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

func main() {
	//webhookUrl := "https://hooks.slack.com/services/X1234"
	webhookUrl := "https://app.slack.com/client/TLNQNBVSS/CMR79D4J1/details/info"
	err := SendSlackNotification(webhookUrl, "Test Message from sillyhat.com")
	if err != nil {
		log.Fatal(err)
	}
}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func SendSlackNotification(webhookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		fmt.Println(buf.String())
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
