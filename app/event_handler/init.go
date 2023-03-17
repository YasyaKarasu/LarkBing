package eventHandler

import (
	receiveMessage "xlab-feishu-robot/app/event_handler/receive_message"
	eventDispatcher "xlab-feishu-robot/pkg/event_dispatcher"
)

func Init() {
	// register your handlers here
	// example
	eventDispatcher.RegisterListener(receiveMessage.Receive, "im.message.receive_v1")
}
