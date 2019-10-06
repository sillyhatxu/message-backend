package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sillyhatxu/http-client"
	"github.com/sillyhatxu/message-backend/slack/dto"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Client struct {
	host string
	url  string
}

func NewSlackClient(host string, url string) *Client {
	return &Client{host: host, url: url}
}

func (c *Client) Send(input dto.SlackDTO) error {
	inputJSON, err := json.Marshal(input)
	if err != nil {
		return err
	}
	httpClient := client.NewHttpClient(c.host)
	response := httpClient.DoPost(c.url, bytes.NewBuffer(inputJSON))
	if response.Error != nil {
		return response.Error
	}
	if response.HttpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("response error. code : %s", response.HttpResponse.Status)
	}
	bodyBytes, err := ioutil.ReadAll(response.HttpResponse.Body)
	if err != nil {
		logrus.Errorf("analysis response body error. %v", err)
		return err
	}
	logrus.Infof("response body : %v", string(bodyBytes))
	return nil
}
