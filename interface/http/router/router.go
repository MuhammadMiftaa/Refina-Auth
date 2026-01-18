package router

import (
	"refina-auth/config/db"
	"refina-auth/config/redis"
	"refina-auth/interface/http/middleware"
	"refina-auth/interface/http/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware(), middleware.GinMiddleware())

	router.GET("test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	routes.UserRoutes(router, db.DB, redis.RDB)

	return router
}
