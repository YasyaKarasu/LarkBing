package receiveButton

import (
	receiveMessage "LarkBing/app/event_handler/receive_message"
	"LarkBing/pkg/bing"
	"context"

	"github.com/sirupsen/logrus"
)

func groupChat(operatorId string, message string) {
	messageevent := receiveMessage.MessageEvent{}
	messageevent.Message.Chat_id = operatorId
	messageevent.Message.Chat_type = "group"
	bingCli := bing.New()
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", &messageevent),
		message,
	)
}
