package http

import (
	"orders/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func routes(h handler.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("/")
	api.GET("ping", h.Ping)
	{
		orderApi := api.Group("order")
		{
			orderApi.GET(":id", h.Order().GetByID)
			orderApi.POST("", h.Order().Create)
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
