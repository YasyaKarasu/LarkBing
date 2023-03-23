package receiveButton

import (
	receiveMessage "LarkBing/app/event_handler/receive_message"
	"LarkBing/pkg/bing"
	"context"

	"github.com/sirupsen/logrus"
)

func p2pChat(operatorId string, message string) {
	messageevent := receiveMessage.MessageEvent{}
	messageevent.Sender.Sender_id.Open_id = operatorId
	messageevent.Message.Chat_type = "p2p"
	bingCli := bing.GetBingClient(operatorId)
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", &messageevent),
		message,
	)
}
