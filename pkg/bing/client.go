package bing

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"xlab-feishu-robot/config"
	"xlab-feishu-robot/pkg/session"

	"github.com/sirupsen/logrus"
)

type BingClient struct {
	ConversationID        string `json:"conversationId"`
	ConversationSignature string `json:"conversationSignature"`
	ClientID              string `json:"clientId"`
	IsStartofSession      bool   `json:"isStartOfSession"`
}

func New() *BingClient {
	req, err := http.NewRequest("GET", "https://www.bing.com/turing/conversation/create", nil)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	req.AddCookie(&http.Cookie{Name: "_U", Value: config.C.Bing.Cookie})

	cli := &http.Client{
		Timeout: time.Second * 15,
	}

	resp, err := cli.Do(req)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	return &BingClient{
		ConversationID:        result["conversationId"].(string),
		ConversationSignature: result["conversationSignature"].(string),
		ClientID:              result["clientId"].(string),
		IsStartofSession:      true,
	}
}

func GetBingClient(ID string) *BingClient {
	clientRaw := session.GetSession(ID)
	if clientRaw == "" {
		client := New()
		bytes, _ := json.Marshal(client)
		session.SetSession(ID, string(bytes))
		return client
	} else {
		var client BingClient
		json.Unmarshal([]byte(clientRaw), &client)
		return &client
	}
}
