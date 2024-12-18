package router

import (
	"PhotonTrail-backend/internal/controller"
	"PhotonTrail-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initStsRouters(r *gin.RouterGroup) {
	sts := r.Group("/sts")
	sts.Use(middleware.JwtAuthMiddleware())
	sts.GET("", controller.GetStsToken) // get sts token
}
