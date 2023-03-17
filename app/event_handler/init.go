package eventHandler

import (
	receiveMessage "xlab-feishu-robot/app/event_handler/receive_message"
	"xlab-feishu-robot/pkg/eventDispatcher"
)

func Init() {
	// register your handlers here
	// example
	eventDispatcher.RegisterListener(receiveMessage.Receive, "im.message.receive_v1")
}
