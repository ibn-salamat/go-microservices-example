package http

import (
	"users/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func routes(h handler.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("/")
	api.GET("ping", h.Ping)
	{
		userApi := api.Group("/user")
		{
			userApi.GET("", h.User().Get)
		}
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
