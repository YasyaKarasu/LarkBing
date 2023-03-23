package bing

import (
	"LarkBing/pkg/session"
	"context"
	"encoding/json"
	"strconv"

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
	Source                string          `json:"source"`
	OptionsSets           []string        `json:"optionsSets"`
	AllowedMessageTypes   []string        `json:"allowedMessageTypes"`
	IsStartofSession      bool            `json:"isStartOfSession"`
	Message               Message         `json:"message"`
	ConversationSignature string          `json:"conversationSignature"`
	Participant           ChatParticipant `json:"participant"`
	ConversationID        string          `json:"conversationId"`
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
				Participant:           ChatParticipant{ID: c.ClientID},
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

func (data *MessageData) WithInvocationID(id int) *MessageData {
	data.InvocationID = strconv.Itoa(id)
	return data
}

type ChatItemHandler func(context.Context, Item, bool, int)

var chatItemHandler ChatItemHandler

func RegisterChatResponseHandler(handler ChatItemHandler) {
	chatItemHandler = handler
}

func (c *BingClient) Chat(ctx context.Context, question string) {
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
	wsCli.SetReadLimit(327680)

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

	data := c.DefaultMessageData().WithText(question).WithInvocationID(c.InvocationID)
	logrus.Info("send data: ", *data)
	b, _ := json.Marshal(data)
	err = wsCli.Send(
		context.Background(),
		websocket.MessageText,
		append(b, 0x1e),
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
		msgContent = msgContent[:len(msgContent)-1]
		if msgType == websocket.MessageText {
			msgData := tools.Any2json(msgContent)
			switch msgData.Get("type").Int() {
			case 1:
				var chatResponse ChatResponse
				type ChatUpdateArgument struct {
					Messages  []ChatMessage `json:"messages"`
					RequestID string        `json:"requestId"`
				}
				type ChatUpdate struct {
					Type      int                  `json:"type"`
					Target    string               `json:"target"`
					Arguments []ChatUpdateArgument `json:"arguments"`
				}
				var chatUpdate ChatUpdate
				json.Unmarshal(msgContent, &chatUpdate)
				if len(chatUpdate.Arguments) == 0 {
					continue
				}
				storedMessages := make([]ChatMessage, 0)
				json.Unmarshal(
					[]byte(session.GetSessionString("messages_"+chatUpdate.Arguments[0].RequestID)),
					&storedMessages,
				)
				if len(storedMessages) == 0 {
					chatResponse.Item = Item{
						Messages:  chatUpdate.Arguments[0].Messages,
						RequestID: chatUpdate.Arguments[0].RequestID,
					}
				} else {
					chatResponse.Item = Item{
						Messages:  storedMessages,
						RequestID: chatUpdate.Arguments[0].RequestID,
					}
					for _, newMsg := range chatUpdate.Arguments[0].Messages {
						for j, oldMsg := range storedMessages {
							if newMsg.MessageID == oldMsg.MessageID {
								chatResponse.Item.Messages[j] = newMsg
								break
							} else {
								if j == len(storedMessages)-1 {
									chatResponse.Item.Messages = append(chatResponse.Item.Messages, newMsg)
								}
							}
						}
					}
				}
				bytes, _ := json.Marshal(chatResponse.Item.Messages)
				session.SetSession("messages_"+chatUpdate.Arguments[0].RequestID, string(bytes))
				chatItemHandler(ctx, chatResponse.Item, true, c.InvocationID)
			case 2:
				run = false
				var chatResponse ChatResponse
				err := json.Unmarshal(msgContent, &chatResponse)
				if err != nil {
					json.Unmarshal(msgContent[:len(msgContent)-30], &chatResponse)
					// len("\x1e{\"type\":3,\"invocationId\":\"1\"}") = 30
				}
				chatItemHandler(ctx, chatResponse.Item, false, c.InvocationID)
			case 7:
				if msgData.Get("allowReconnect").Bool() {
					err = wsCli.Send(
						context.Background(),
						websocket.MessageText,
						append(tools.StringToBytes(`{"type":6}`), 0x1e),
					)
					if err != nil {
						logrus.Error(err)
						return
					}
				}
			default:
				logrus.Info(string(msgContent))
			}
		}
	}
}
