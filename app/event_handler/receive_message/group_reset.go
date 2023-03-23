package receiveMessage

import (
	"LarkBing/pkg/global"
	"LarkBing/pkg/session"

	"github.com/YasyaKarasu/feishuapi"
)

func init() {
	groupMessageRegister(groupReset, "/重置")
}

func groupReset(messageevent *MessageEvent) {
	session.ClearSession(messageevent.Message.Chat_id)
	global.Cli.MessageSend(feishuapi.GroupChatId, messageevent.Message.Chat_id, feishuapi.Text, "重置成功！")
}
