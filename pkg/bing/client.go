package bing

import (
	"context"
	"encoding/json"
	"net/http"
	"xlab-feishu-robot/config"
	"xlab-feishu-robot/pkg/session"

	"gitee.com/baixudong/gospider/requests"
	"github.com/sirupsen/logrus"
)

type BingClient struct {
	ConversationID        string `json:"conversationId"`
	ConversationSignature string `json:"conversationSignature"`
	ClientID              string `json:"clientId"`
	IsStartofSession      bool   `json:"isStartOfSession"`
}

func New() *BingClient {
	reqCli, err := requests.NewClient(context.Background())
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer reqCli.Close()

	response, err := reqCli.Request(context.Background(), "get", "https://www.bing.com/turing/conversation/create",
		requests.RequestOption{
			Cookies: http.Cookie{Name: "_U", Value: config.C.Bing.Cookie},
		},
	)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	jsonData := response.Json()
	return &BingClient{
		ConversationID:        jsonData.Get("conversationId").String(),
		ConversationSignature: jsonData.Get("conversationSignature").String(),
		ClientID:              jsonData.Get("clientId").String(),
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
