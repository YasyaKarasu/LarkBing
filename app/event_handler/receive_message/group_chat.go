package receiveMessage

import (
	"LarkBing/pkg/bing"
	"context"

	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

func groupChat(messageevent *MessageEvent) {
	bingCli := bing.New()
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	b, _ := json.Marshal(messageevent)
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", b),
		messageevent.Message.Content,
	)
}
