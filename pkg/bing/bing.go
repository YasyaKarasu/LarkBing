package bing

import (
	"context"
	"encoding/json"

	"gitee.com/baixudong/gospider/requests"
	"gitee.com/baixudong/gospider/tools"
	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
)

type Message struct {
	Text        string `json:"text"`
	MessageType string `json:"messageType"`
}

type ChatParticipant struct {
	ID string `json:"id"`
}

type MessageArgument struct {
	Source                string            `json:"source"`
	OptionsSets           []string          `json:"optionsSets"`
	AllowedMessageTypes   []string          `json:"allowedMessageTypes"`
	IsStartofSession      bool              `json:"isStartOfSession"`
	Message               Message           `json:"message"`
	ConversationSignature string            `json:"conversationSignature"`
	Participant           []ChatParticipant `json:"participant"`
	ConversationID        string            `json:"conversationId"`
}

type MessageData struct {
	Arguments    []MessageArgument `json:"arguments"`
	InvocationID string            `json:"invocationId"`
	Target       string            `json:"target"`
	Type         int               `json:"type"`
}

func (c *BingClient) DefaultMessageData() *MessageData {
	return &MessageData{
		Arguments: []MessageArgument{
			{
				Source: "cib",
				OptionsSets: []string{
					"nlu_direct_response_filter",
					"deepleo",
					"enable_debug_commands",
					"disable_emoji_spoken_text",
					"responsible_ai_policy_235",
					"enablemm",
				},
				AllowedMessageTypes: []string{
					"Chat",
					"InternalSearchQuery",
					"InternalSearchResult",
					"Disengaged",
					"InternalLoaderMessage",
					"RenderCardRequest",
					"AdsQuery",
					"SemanticSerp",
					"GenerateContentQuery",
					"SearchQuery",
				},
				IsStartofSession:      c.IsStartofSession,
				Message:               Message{Text: "", MessageType: "Chat"},
				ConversationSignature: c.ConversationSignature,
				Participant:           []ChatParticipant{{ID: c.ClientID}},
				ConversationID:        c.ConversationID,
			},
		},
		InvocationID: "1",
		Target:       "chat",
		Type:         4,
	}
}

func (data *MessageData) WithText(text string) *MessageData {
	data.Arguments[0].Message.Text = text
	return data
}

func (c *BingClient) Chat(question string) {
	reqCli, err := requests.NewClient(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	defer reqCli.Close()

	response, err := reqCli.Request(context.Background(), "get", "wss://sydney.bing.com/sydney/ChatHub")
	if err != nil {
		logrus.Error(err)
		return
	}

	wsCli := response.WebSocket()

	err = wsCli.Send(
		context.Background(),
		websocket.MessageText,
		append(tools.StringToBytes(`{"protocol":"json","version":1}`), 0x1e),
	)
	if err != nil {
		logrus.Error(err)
		return
	}

	wsCli.Recv(context.Background())

	err = wsCli.Send(
		context.Background(),
		websocket.MessageText,
		append(tools.StringToBytes(`{"type":6}`), 0x1e),
	)
	if err != nil {
		logrus.Error(err)
		return
	}

	data := make(map[string]any)
	struct2map(c.DefaultMessageData().WithText(question), &data)
	logrus.Info("send data: ", data)
	err = wsCli.Send(
		context.Background(),
		websocket.MessageText,
		append(tools.StringToBytes(tools.Any2json(data).Raw), 0x1e),
	)
	if err != nil {
		logrus.Error(err)
		return
	}

	run := true
	for run {
		msgType, msgContent, err := wsCli.Recv(context.Background())
		if err != nil {
			logrus.Error(err)
			break
		}
		if msgType == websocket.MessageText {
			msgData := tools.Any2json(msgContent)
			logrus.Info("receive msgData: ", msgData)
			switch msgData.Get("type").Int() {
			case 1:
				text := msgData.Get("arguments.0.messages.text").String()
				runeText := []rune(text)
				logrus.Info(runeText)
			case 2:
				logrus.Info(msgData)
				run = false
			}
		}
	}
}

func struct2map(s any, m any) {
	b, _ := json.Marshal(s)
	json.Unmarshal(b, m)
}
