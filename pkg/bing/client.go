package bing

import (
	"context"
	"encoding/json"
	"log"
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
		log.Panic(err)
	}
	response, err := reqCli.Request(context.Background(), "get", "https://www.bing.com/turing/conversation/create", requests.RequestOption{
		Cookies: "_U=" + config.C.Bing.Cookie,
	})
	if err != nil {
		log.Panic(err)
	}
	jsonData := response.Json()

	logrus.Info(jsonData)

	conversationId := jsonData.Get("conversationId").String()
	clientId := jsonData.Get("clientId").String()
	conversationSignature := jsonData.Get("conversationSignature").String()

	return &BingClient{
		ConversationID:        conversationId,
		ClientID:              clientId,
		ConversationSignature: conversationSignature,
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
