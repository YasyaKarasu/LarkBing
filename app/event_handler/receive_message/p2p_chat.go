package receiveMessage

import (
	"LarkBing/pkg/bing"

	"github.com/sirupsen/logrus"
)

func p2pChat(messageevent *MessageEvent) {
	bingCli := bing.GetBingClient(messageevent.Sender.Sender_id.Open_id)
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	bingCli.Chat(messageevent.Message.Content)
}
