package controller

import (
	"PhotonTrail-backend/internal/global"
	"PhotonTrail-backend/internal/global/param"
	"PhotonTrail-backend/internal/global/response"
	"PhotonTrail-backend/internal/model"
	"PhotonTrail-backend/pkg/common"
	"PhotonTrail-backend/pkg/jwt"
	"PhotonTrail-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	var req param.ReqRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusOK, response.CodeInvalidParams, "邮箱或密码格式错误", err.Error())
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "hash password failed", err.Error())
		return
	}
	user := &model.User{
		Name:        util.GenerateRandomName(),
		Email:       req.Email,
		Password:    hashedPassword,
		Role:        common.User,
		Avatar:      "",
		Description: "",
		Background:  "",
	}

	id, err := user.Create(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "create user failed", err.Error())
		return
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"id": id}, "注册成功")
}

func Login(c *gin.Context) {
	var req param.ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusOK, response.CodeInvalidParams, "邮箱或密码格式错误", err.Error())
		return
	}

	user := &model.User{
		Email: req.Email,
	}
	user, err := user.GetByEmail(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get user failed", err.Error())
		return
	}

	if ok := util.CheckHashedPassword(req.Password, user.Password); !ok {
		response.Error(c, http.StatusOK, response.CodeInvalidParams, "邮箱或密码错误", "")
		return
	}

	tokenClaims, expireTime := jwt.GenerateJwtToken(strconv.Itoa(int(user.ID)), user.Role, global.Config.Jwt.Expire, global.Config.Jwt.Issuer)
	token, err := jwt.GenerateJwtTokenString(tokenClaims, []byte(global.Config.Jwt.Key))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "generate jwt token failed", err.Error())
		return
	}

	userData := response.UserData{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		Avatar:      user.Avatar,
		Description: user.Description,
		Background:  user.Background,
	}

	response.Success(
		c, http.StatusOK,
		response.CodeSuccess,
		response.Data{
			"id":      user.ID,
			"role":    user.Role,
			"token":   token,
			"expire":  expireTime,
			"profile": userData,
		},
		"登录成功")
}

func GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}
	user := &model.User{
		Model: &gorm.Model{ID: uint(idUint)},
	}
	user, err = user.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get user failed", err.Error())
		return
	}

	userData := response.UserData{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		Avatar:      user.Avatar,
		Description: user.Description,
		Background:  user.Background,
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"profile": userData}, "get user profile success")
}

func UpdateUserProfile(c *gin.Context) {
	userData, _ := c.Get("user")
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}
	if idInt != uidInt {
		response.Error(c, http.StatusForbidden, response.CodeInvalidParams, "forbidden", "")
		return
	}

	var req param.ReqUpdateProfile
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusOK, response.CodeInvalidParams, "资料格式错误", err.Error())
		return
	}

	user := &model.User{
		Model: &gorm.Model{ID: uint(uidInt)},
	}
	err = user.Update(global.DBEngine, map[string]interface{}{
		"name":        req.Name,
		"avatar":      req.Avatar,
		"description": req.Description,
		"background":  req.Background,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "update user profile failed", "")
		return
	}

	user, err = user.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "get user failed", err.Error())
		return
	}
	profile := response.UserData{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		Avatar:      user.Avatar,
		Description: user.Description,
		Background:  user.Background,
	}
	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"profile": profile}, "资料更新成功")
}
