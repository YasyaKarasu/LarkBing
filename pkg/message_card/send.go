package messageCard

import (
	receiveMessage "LarkBing/app/event_handler/receive_message"
	"LarkBing/pkg/bing"
	"LarkBing/pkg/global"
	messageCardDispatcher "LarkBing/pkg/message_card_dispatcher"
	"LarkBing/pkg/session"
	"context"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/YasyaKarasu/feishuapi"
	"github.com/sirupsen/logrus"
)

func generateCard(item bing.Item, updating bool) string {
	type element struct {
		tag  string
		text string
	}

	elements := make([]element, 0)
	var referenceItems []bing.SourceAttribution
	var suggestedItems []bing.SuggestedResponse

	for _, message := range item.Messages {
		if message.MessageType == "InternalSearchResult" || message.MessageType == "RenderCardRequest" || message.Text == "" {
			continue
		} else if message.MessageType != "" {
			if message.MessageType == "InternalSearchQuery" {
				elements = append(elements, element{
					tag:  "Note",
					text: "🔍 " + strings.ReplaceAll(message.Text, "`", " "),
				})
			} else if message.MessageType == "InternalLoaderMessage" {
				elements = append(elements, element{
					tag:  "Note",
					text: "🤖️ " + message.Text,
				})
			} else {
				elements = append(elements, element{
					tag:  "Note",
					text: "💡 " + message.Text,
				})
			}
		} else {
			if message.Author == "user" {
				elements = append(elements, element{
					tag:  "Note",
					text: "🗣️ Question: " + message.Text,
				})
			} else if message.Author == "bot" {
				text := message.Text
				referenceItems = message.SourceAttributions
				suggestedItems = message.SuggestedResponses

				reg := regexp.MustCompile(`\[\^[0-9]+\^\]`)
				for index, val := range reg.FindAllString(message.Text, -1) {
					if len(referenceItems) >= index {
						text = strings.Replace(text,
							val,
							"[["+strconv.Itoa(index)+"]]("+referenceItems[index-1].SeeMoreURL+")",
							-1,
						)
					} else {
						text = strings.Replace(text, val, "["+strconv.Itoa(index)+"]", -1)
					}
				}

				reg = regexp.MustCompile(`\x60\x60\x60(\w+)\n`)
				reg.ReplaceAllString(text, "  💾 Code:\n━━━━━━━━━━━━\n")

				text = strings.ReplaceAll(text, "```", "\n━━━━━━━━━━━━")

				elements = append(elements, element{
					tag:  "Answer",
					text: text,
				})
			}
		}
	}

	type Config struct {
		WideScreenMode bool `json:"wide_screen_mode"`
		UpdateMulti    bool `json:"update_multi"`
	}

	type Title struct {
		Tag     string `json:"tag,omitempty"`
		Content string `json:"content,omitempty"`
	}

	type Header struct {
		Title    Title  `json:"title,omitempty"`
		Template string `json:"template,omitempty"`
	}

	type CardTemplate struct {
		Config   Config `json:"config"`
		Header   Header `json:"header,omitempty"`
		Elements []any  `json:"elements"`
	}

	result := CardTemplate{
		Config: Config{
			WideScreenMode: true,
			UpdateMulti:    true,
		},
		Header: Header{
			Title: Title{
				Tag:     "plain_text",
				Content: "⚙️ Updating...",
			},
			Template: "blue",
		},
		Elements: make([]any, 0),
	}

	if !updating {
		result.Header = Header{}
	}

	for _, val := range elements {
		if val.tag == "Note" {
			type NodeElement struct {
				Tag     string `json:"tag"`
				Content string `json:"content"`
			}
			type Note struct {
				Tag      string        `json:"tag"`
				Elements []NodeElement `json:"elements"`
			}
			result.Elements = append(result.Elements, Note{
				Tag: "note",
				Elements: []NodeElement{{
					Tag:     "lark_md",
					Content: val.text,
				}},
			})
		} else if val.tag == "Answer" {
			if len(result.Elements) > 0 {
				result.Elements = append(result.Elements, struct {
					Tag string `json:"tag"`
				}{Tag: "hr"})
			}
			result.Elements = append(result.Elements, struct {
				Tag     string `json:"tag"`
				Content string `json:"content"`
			}{Tag: "markdown", Content: val.text})
		}
	}

	if len(referenceItems) > 0 {
		result.Elements = append(result.Elements, struct {
			Tag string `json:"tag"`
		}{Tag: "hr"})

		type Text struct {
			Tag     string `json:"tag"`
			Content string `json:"content"`
		}
		type ReferenceOption struct {
			Text  Text   `json:"text"`
			Value string `json:"value"`
			URL   string `json:"url"`
		}
		referenceOptions := make([]ReferenceOption, 0)
		for index, value := range referenceItems {
			referenceOptions = append(referenceOptions, ReferenceOption{
				Text: Text{
					Tag:     "plain_text",
					Content: "[" + strconv.Itoa(index+1) + "] " + value.ProviderDisplayName,
				},
				Value: value.SeeMoreURL,
				URL:   value.SeeMoreURL,
			})
		}

		type DivExtra struct {
			Tag     string            `json:"tag"`
			Options []ReferenceOption `json:"options"`
		}
		result.Elements = append(result.Elements, struct {
			Tag   string   `json:"tag"`
			Text  Text     `json:"text"`
			Extra DivExtra `json:"extra"`
		}{
			Tag: "div",
			Text: Text{
				Tag:     "lark_md",
				Content: "  Learn more / Reference 👉",
			},
			Extra: DivExtra{
				Tag:     "overflow",
				Options: referenceOptions,
			},
		})
	}

	if len(suggestedItems) > 0 {
		type ButtonValue struct {
			Text string `json:"text"`
		}
		type Text struct {
			Tag     string `json:"tag"`
			Content string `json:"content"`
		}
		type Button struct {
			Tag   string      `json:"tag"`
			Value ButtonValue `json:"value"`
			Text  Text        `json:"text"`
			Type  string      `json:"type"`
		}
		type ButtonInteract struct {
			Tag     string   `json:"tag"`
			Actions []Button `json:"actions"`
		}

		buttons := make([]Button, 0)
		for _, value := range suggestedItems {
			buttons = append(buttons, Button{
				Tag: "button",
				Value: ButtonValue{
					Text: value.Text,
				},
				Text: Text{
					Tag:     "plain_text",
					Content: value.Text,
				},
				Type: "primary",
			})
		}
		result.Elements = append(result.Elements, ButtonInteract{
			Tag:     "action",
			Actions: buttons,
		})
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return string(bytes)
}

func SendCard(ctx context.Context, item bing.Item, updating bool) {
	card := generateCard(item, updating)
	if mid := session.GetSessionString(item.RequestID); mid != "" {
		global.Cli.UpdateMessage(mid, card)
	} else {
		messageevent := ctx.Value("messageevent").(receiveMessage.MessageEvent)
		var mid string
		var status messageCardDispatcher.MessageCardState
		switch messageevent.Message.Chat_type {
		case "p2p":
			mid, _ = global.Cli.MessageSend(
				feishuapi.UserOpenId,
				messageevent.Sender.Sender_id.Open_id,
				feishuapi.Interactive,
				card,
			)
			status = messageCardDispatcher.MessageCardState{
				OperatorId: messageevent.Sender.Sender_id.Open_id,
				ChatType:   "p2p",
			}
		case "group":
			mid, _ = global.Cli.MessageSend(
				feishuapi.GroupChatId,
				messageevent.Message.Chat_id,
				feishuapi.Interactive,
				card,
			)
			status = messageCardDispatcher.MessageCardState{
				OperatorId: messageevent.Message.Chat_id,
				ChatType:   "group",
			}
		}
		session.SetSession(item.RequestID, mid)
		bytes, _ := json.Marshal(status)
		session.SetSession(mid, string(bytes))
	}
}
