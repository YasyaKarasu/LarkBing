package router

import (
	"LarkBing/app/controller"
	eventDispatcher "LarkBing/pkg/event_dispatcher"
	messageCardDispatcher "LarkBing/pkg/message_card_dispatcher"

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
