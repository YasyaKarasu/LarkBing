package receiveButton

import (
	"encoding/json"
	"xlab-feishu-robot/pkg/session"
)

type messageCardState struct {
	OperatorId string
	ChatType   string
}

func Receive(OpenMessageId string, action map[string]any) {
	value := action["value"].(map[string]any)
	var messageState messageCardState
	json.Unmarshal([]byte(session.GetSession(OpenMessageId)), &messageState)
	send(OpenMessageId, value, messageState)
}
