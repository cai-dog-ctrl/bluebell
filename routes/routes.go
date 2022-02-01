package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v2 := r.Group("/api/v1")
	v2.POST("/signup", controllers.SingUpHandler)
	v2.POST("/login", controllers.LoginHandler)
	v2.Use(middleware.JWTAuthMiddleware())
	{
		v2.GET("/community", controllers.CommunityHandler)
		v2.GET("/community/:id", controllers.CommunityDetailHandler)
		v2.POST("/post", controllers.CreatPost)
		v2.GET("/post/:id", controllers.GetPostById)
		v2.GET("/posts", controllers.GetPostList)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
