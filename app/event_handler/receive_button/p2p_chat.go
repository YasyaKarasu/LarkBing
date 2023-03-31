package receiveButton

import (
	receiveMessage "LarkBing/app/event_handler/receive_message"
	"LarkBing/pkg/bing"
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func p2pChat(operatorId string, message string) {
	messageevent := receiveMessage.MessageEvent{}
	messageevent.Message.Chat_id = operatorId
	messageevent.Message.Chat_type = "p2p"
	bingCli := bing.New()
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	b, _ := json.Marshal(messageevent)
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", b),
		message,
	)
}
