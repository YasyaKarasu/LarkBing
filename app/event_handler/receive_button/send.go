package receiveButton

import (
	"xlab-feishu-robot/pkg/global"
	messageCardDispatcher "xlab-feishu-robot/pkg/message_card_dispatcher"

	"github.com/YasyaKarasu/feishuapi"
)

func send(OpenMessageId string, value map[string]string, MessageState messageCardDispatcher.MessageCardState) {
	content := value["content"]
	if MessageState.ChatType == "p2p" {
		global.Cli.MessageSend(feishuapi.UserOpenId, MessageState.OperatorId, feishuapi.Text, content)
	} else {
		global.Cli.MessageSend(feishuapi.GroupChatId, MessageState.OperatorId, feishuapi.Text, content)
	}
}
