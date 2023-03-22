package receiveMessage

import (
	"LarkBing/pkg/global"

	_ "github.com/sirupsen/logrus"
)

func init() {
	groupMessageRegister(groupHelpMenu, "help")
}

func groupHelpMenu(messageevent *MessageEvent) {
	global.Cli.MessageSend("chat_id", messageevent.Message.Chat_id, "text", "this is a group test string")
}
