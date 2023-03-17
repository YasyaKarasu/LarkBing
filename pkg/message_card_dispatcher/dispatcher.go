package message_card_dispatcher

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Dispatcher(c *gin.Context) {
	rawBody, _ := io.ReadAll(c.Request.Body)
	requestStr := string(rawBody)

	var req FeishuMessageCardRequest
	deserializeRequest(requestStr, &req)
	logrus.Debug("Feishu robot received a request: ", req)

	if req.Challenge != "" {
		c.JSON(http.StatusOK, gin.H{"challenge": req.Challenge})
		return
	}
}
