package receiveMessage

import (
	"LarkBing/pkg/global"
	"LarkBing/pkg/session"

	"github.com/YasyaKarasu/feishuapi"
)

func init() {
	p2pMessageRegister(p2pReset, "/重置")
}

func p2pReset(messageevent *MessageEvent) {
	session.ClearSession("jailbreak_" + messageevent.Message.Chat_id)
	global.Cli.MessageSend(feishuapi.UserOpenId, messageevent.Sender.Sender_id.Open_id, feishuapi.Text, "重置成功！")
}
