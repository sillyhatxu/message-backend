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

curl -X POST -H 'Authorization: Bearer xoxb-1234-56789abcdefghijklmnop' \
-H 'Content-type: application/json' \
--data '{"channel":"C061EG9SL","text":"I hope the tour went well, Mr. Wonka.","attachments": [{"text":"Who wins the lifetime supply of chocolate?","fallback":"You could be telling the computer exactly what it can do with a lifetime supply of chocolate.","color":"#3AA3E3","attachment_type":"default","callback_id":"select_simple_1234","actions":[{"name":"winners_list","text":"Who should win?","type":"select","data_source":"users"}]}]}' \
https://slack.com/api/chat.postMessage

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
