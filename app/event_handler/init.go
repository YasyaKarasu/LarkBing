package eventHandler

import (
	receiveMessage "LarkBing/app/event_handler/receive_message"
	eventDispatcher "LarkBing/pkg/event_dispatcher"
)

func Init() {
	// register your handlers here
	// example
	eventDispatcher.RegisterListener(receiveMessage.Receive, "im.message.receive_v1")
}
