package receiveMessage

import (
	"LarkBing/pkg/bing"
	"context"

	"github.com/sirupsen/logrus"
)

func groupChat(messageevent *MessageEvent) {
	bingCli := bing.New()
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", messageevent),
		messageevent.Message.Content,
	)
}
