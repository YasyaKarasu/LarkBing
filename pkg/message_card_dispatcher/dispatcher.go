package messageCardDispatcher

import (
	"io"
	"net/http"
	"xlab-feishu-robot/pkg/session"

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

	if requestRepeatDetect(c) {
		logrus.Warn("Repeated message card request: ", req)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "repeated message card request"})
		return
	}

	if req.Action["tag"].(string) != "button" {
		logrus.WithField("tag", req.Action["tag"].(string)).Warn("Received unsupported message card event")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "unsupported message card event"})
		return
	}

	// dispatch events
}

func requestRepeatDetect(c *gin.Context) bool {
	refreshToken := c.Request.Header.Get("X-Refresh-Token")
	storedToken := session.GetSession(refreshToken)
	if storedToken == "" {
		session.SetSession(refreshToken, "true")
		return false
	} else {
		return true
	}
}
