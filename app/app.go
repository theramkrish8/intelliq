package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	"intelliq/app/approuter"
	"intelliq/app/config"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	approuter.AddRouters(router)
	config.Connect()
	//router.Use(cors.Default())
	enableCors()
	enableSession()
	router.Run()
}

func enableCors() {

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "currentrole", "Content-Type", "X-Requested-With", "Accept"},
		ExposeHeaders:          []string{"currentrole"},
		AllowCredentials:       false,
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	}))
}

func enableSession() HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("intelliq_session"))
	router.Use(sessions.Sessions("redisSession", store))
}
