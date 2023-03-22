package receiveButton

import (
	"LarkBing/pkg/global"

	"github.com/YasyaKarasu/feishuapi"
)

func send(OpenMessageId string, value map[string]any, MessageState messageCardState) {
	content := value["content"].(string)
	if MessageState.ChatType == "p2p" {
		global.Cli.MessageSend(feishuapi.UserOpenId, MessageState.OperatorId, feishuapi.Text, content)
	} else {
		global.Cli.MessageSend(feishuapi.GroupChatId, MessageState.OperatorId, feishuapi.Text, content)
	}
}
