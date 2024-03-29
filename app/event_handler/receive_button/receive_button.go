package receiveButton

import (
	"LarkBing/pkg/session"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type messageCardState struct {
	OperatorId string
	ChatType   string
}

func Receive(OpenMessageId string, action map[string]any) {
	logrus.Info(action)
	value := action["value"].(map[string]any)
	var messageState messageCardState
	json.Unmarshal([]byte(session.GetSessionString(OpenMessageId)), &messageState)
	if messageState.ChatType == "p2p" {
		p2pChat(messageState.OperatorId, value["text"].(string))
	} else {
		groupChat(messageState.OperatorId, value["text"].(string))
	}
}
