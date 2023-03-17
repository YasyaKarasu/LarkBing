package receiveButton

import (
	"encoding/json"
	messageCardDispatcher "xlab-feishu-robot/pkg/message_card_dispatcher"
	"xlab-feishu-robot/pkg/session"
)

func Receive(OpenMessageId string, action map[string]any) {
	value := action["value"].(map[string]string)
	var messageState messageCardDispatcher.MessageCardState
	json.Unmarshal([]byte(session.GetSession(OpenMessageId)), &messageState)
	send(OpenMessageId, value, messageState)
}
