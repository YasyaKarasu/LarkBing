package router

import (
	"xlab-feishu-robot/app/controller"
	"xlab-feishu-robot/pkg/event_dispatcher"
	"xlab-feishu-robot/pkg/message_card_dispatcher"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// register your controllers here
	// example
	r.POST("/api/example", controller.Example)

	// DO NOT CHANGE LINES BELOW
	// register dispatcher
	r.POST("/feiShu/Event", event_dispatcher.Dispatcher)

	r.POST("/feiShu/MessageCard", message_card_dispatcher.Dispatcher)
}
