package response

import (
	"PhotonTrail-backend/internal/global"
	"PhotonTrail-backend/pkg/common"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code  ErrorCode `json:"code"`
	Data  Data      `json:"data,omitempty"`
	Msg   string    `json:"msg,omitempty"`
	Error string    `json:"error,omitempty"` // only available in debug mode
}

type Data map[string]any

type UserData struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Role        common.Role `json:"role"`
	Avatar      string      `json:"avatar"`
	Description string      `json:"description"`
	Background  string      `json:"background"`
}

type StsData struct {
	Dir                  string `json:"dir"`
	Host                 string `json:"host"`
	Policy               string `json:"policy"`
	SecurityToken        string `json:"security_token"`
	Signature            string `json:"signature"`
	XOssCredential       string `json:"x_oss_credential"`
	XOssDate             string `json:"x_oss_date"`
	XOssSignatureVersion string `json:"x_oss_signature_version"`
}

type PostData struct {
	ID          uint        `json:"id"`
	UserID      uint        `json:"user_id"`
	Avatar      string      `json:"avatar"`
	Name        string      `json:"name"`
	PublishedAt string      `json:"published_at"`
	Images      []PostImage `json:"images"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	Camera      string      `json:"camera"`
	Lens        string      `json:"lens"`
}

type PostImage struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Index  int    `json:"index"`
}

func Success(c *gin.Context, status int, code ErrorCode, data Data, msg string) {
	c.JSON(status, response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Error(c *gin.Context, status int, code ErrorCode, msg string, err ...string) {
	if global.Config.App.Debug {
		c.JSON(status, response{
			Code:  code,
			Msg:   msg,
			Error: err[0],
		})
	} else {
		c.JSON(status, response{
			Code: code,
			Msg:  msg,
		})
	}
}
