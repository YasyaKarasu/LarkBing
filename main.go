package main

import (
	"LarkBing/app"
	"LarkBing/config"
	"LarkBing/docs"
	"LarkBing/pkg/global"
	"LarkBing/pkg/session"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ReadConfig()

	// log
	config.SetupLogrus()
	logrus.Info("Robot starts up")

	// feishu api client
	config.SetupFeishuApiClient(&global.Cli)
	global.Cli.StartTokenTimer()

	session.ConnectRedis()

	// robot server
	r := gin.Default()
	app.Init(r)

	// api docs by swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + fmt.Sprint(config.C.Server.Port))
}
