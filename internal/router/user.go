package router

import (
	"PhotonTrail-backend/internal/controller"
	"PhotonTrail-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.RouterGroup) {
	users := r.Group("/user")
	users.POST("", controller.Register) // register

	userAuth := users.Group("/:id")
	userAuth.Use(middleware.JwtAuthMiddleware())
	userAuth.GET("/profile", controller.GetUserProfile)    // get user profile
	userAuth.PUT("/profile", controller.UpdateUserProfile) // update user profile
	userAuth.GET("/posts", controller.GetUserPosts)        // get user posts

	session := r.Group("/session")
	session.POST("", controller.Login) // login
}
