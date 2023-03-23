package receiveMessage

import (
	"LarkBing/pkg/bing"
	"LarkBing/pkg/global"
	"LarkBing/pkg/session"
	"context"

	"github.com/YasyaKarasu/feishuapi"
	"github.com/sirupsen/logrus"
)

func p2pChat(messageevent *MessageEvent) {
	bingCli := bing.GetBingClient(messageevent.Sender.Sender_id.Open_id)
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	if bingCli.InvocationID >= 16 {
		global.Cli.MessageSend(
			feishuapi.UserOpenId,
			messageevent.Sender.Sender_id.Open_id,
			feishuapi.Text,
			"对话长度已达上限，将自动为您重置对话。",
		)
		session.ClearSession(messageevent.Sender.Sender_id.Open_id)
		bingCli = bing.GetBingClient(messageevent.Sender.Sender_id.Open_id)
	}
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", messageevent),
		messageevent.Message.Content,
	)
}
