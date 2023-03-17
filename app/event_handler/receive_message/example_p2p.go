package receiveMessage

import (
	"xlab-feishu-robot/pkg/global"
	messageCardDispatcher "xlab-feishu-robot/pkg/message_card_dispatcher"
	"xlab-feishu-robot/pkg/session"

	"github.com/YasyaKarasu/feishuapi"
	_ "github.com/sirupsen/logrus"
)

func init() {
	p2pMessageRegister(p2pHelpMenu, "help")
}

func p2pHelpMenu(messageevent *MessageEvent) {
	// global.Cli.MessageSend("open_id", messageevent.Sender.Sender_id.Open_id, "text", "this is a P2P test string")
	content := `{
		"config": {
		  "wide_screen_mode": true
		},
		"elements": [
		  {
			"tag": "action",
			"actions": [
			  {
				"tag": "button",
				"text": {
				  "tag": "plain_text",
				  "content": "button 1"
				},
				"type": "primary",
				"value": {
				  "content": "button 1"
				}
			  }
			]
		  },
		  {
			"tag": "action",
			"actions": [
			  {
				"tag": "button",
				"text": {
				  "tag": "plain_text",
				  "content": "button 2"
				},
				"type": "primary",
				"value": {
				  "content": "button 2"
				}
			  }
			]
		  },
		  {
			"tag": "action",
			"actions": [
			  {
				"tag": "button",
				"text": {
				  "tag": "plain_text",
				  "content": "button 3"
				},
				"type": "primary",
				"value": {
				  "content": "button 3"
				}
			  }
			]
		  }
		],
		"header": {
		  "template": "blue",
		  "title": {
			"content": "这里是卡片标题",
			"tag": "plain_text"
		  }
		}
	  }`
	mid, _ := global.Cli.MessageSend(feishuapi.UserOpenId, messageevent.Sender.Sender_id.Open_id, feishuapi.Interactive, content)
	session.SetSession(mid, messageCardDispatcher.MessageCardState{
		OperatorId: messageevent.Sender.Sender_id.Open_id,
		ChatType:   "p2p",
	})
}
