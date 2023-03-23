package receiveMessage

import (
	"LarkBing/pkg/bing"
	"LarkBing/pkg/global"
	"LarkBing/pkg/session"
	"context"

	"github.com/YasyaKarasu/feishuapi"
	"github.com/sirupsen/logrus"
)

func groupChat(messageevent *MessageEvent) {
	bingCli := bing.GetBingClient(messageevent.Message.Chat_id)
	if bingCli == nil {
		logrus.Error("failed to get bing client")
		return
	}
	if bingCli.InvocationID >= 16 {
		global.Cli.MessageSend(
			feishuapi.GroupChatId,
			messageevent.Message.Chat_id,
			feishuapi.Text,
			"对话长度已达上限，将自动为您重置对话。",
		)
		session.ClearSession(messageevent.Message.Chat_id)
		bingCli = bing.GetBingClient(messageevent.Message.Chat_id)
	}
	bingCli.Chat(
		context.WithValue(context.Background(), "messageevent", messageevent),
		messageevent.Message.Content,
	)
}
