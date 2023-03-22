package receiveMessage

import (
	"LarkBing/pkg/bing"
	"context"

	"github.com/sirupsen/logrus"
)

func p2pChat(messageevent *MessageEvent) {
	bingCli := bing.GetBingClient(messageevent.Sender.Sender_id.Open_id)
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", messageevent),
		messageevent.Message.Content,
	)
}
