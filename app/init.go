package app

import (
	eventHandler "LarkBing/app/event_handler"
	"LarkBing/app/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	eventHandler.Init()
	router.Register(r)
}
