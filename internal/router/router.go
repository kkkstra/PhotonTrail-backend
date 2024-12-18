package router

import (
	"PhotonTrail-backend/internal/global"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group(global.Config.App.ApiPrefix)
	initUserRouters(api)
	initStsRouters(api)
	initPostRouters(api)
}
