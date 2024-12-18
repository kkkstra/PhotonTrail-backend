package router

import (
	"PhotonTrail-backend/internal/controller"
	"PhotonTrail-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initPostRouters(r *gin.RouterGroup) {
	posts := r.Group("/post")
	posts.Use(middleware.JwtAuthMiddleware())

	posts.GET("", controller.GetPosts)
	posts.GET("/:id", controller.GetPost)
	posts.POST("", controller.CreatePost)
	posts.DELETE("/:id", controller.DeletePost)
}
