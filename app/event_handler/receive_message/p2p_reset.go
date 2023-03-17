package receiveMessage

import (
	"xlab-feishu-robot/pkg/global"
	"xlab-feishu-robot/pkg/session"

	"github.com/YasyaKarasu/feishuapi"
)

func init() {
	p2pMessageRegister(p2pReset, "/重置")
}

func p2pReset(messageevent *MessageEvent) {
	session.ClearSession(messageevent.Sender.Sender_id.Open_id)
	global.Cli.MessageSend(feishuapi.UserOpenId, messageevent.Sender.Sender_id.Open_id, feishuapi.Text, "重置成功！")
}
