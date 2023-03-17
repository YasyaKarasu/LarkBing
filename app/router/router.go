package router

import (
	"xlab-feishu-robot/app/controller"
	eventDispatcher "xlab-feishu-robot/pkg/event_dispatcher"
	messageCardDispatcher "xlab-feishu-robot/pkg/message_card_dispatcher"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// register your controllers here
	// example
	r.POST("/api/example", controller.Example)

	// DO NOT CHANGE LINES BELOW
	// register dispatcher
	r.POST("/feiShu/Event", eventDispatcher.Dispatcher)

	r.POST("/feiShu/MessageCard", messageCardDispatcher.Dispatcher)
}
