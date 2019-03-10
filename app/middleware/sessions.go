package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CheckSessionExists checks the session id in header
func CheckSessionExists(c *gin.Context) bool {
	session := sessions.Default(c)
	if len(c.Request.Header.Get("SessionId")) != 0 {
		SessionID := session.Get("SessionId")
		if SessionID != nil {
			return true
		}
		return false
	}
	return false
}
