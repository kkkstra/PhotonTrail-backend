package controller

import (
	"PhotonTrail-backend/internal/global"
	"PhotonTrail-backend/internal/global/param"
	"PhotonTrail-backend/internal/global/response"
	"PhotonTrail-backend/internal/model"
	"PhotonTrail-backend/pkg/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context) {
	userData, _ := c.Get("user")
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)

	var req param.ReqCreatePost
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusOK, response.CodeInvalidParams, "参数错误", err.Error())
		return
	}

	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		Camera:  req.Camera,
		Lens:    req.Lens,
		UserID:  uint(uidInt),
	}

	var images []model.PostImage
	for _, img := range req.Images {
		images = append(images, model.PostImage{
			Url:    img.Url,
			Width:  img.Width,
			Height: img.Height,
			Index:  img.Index,
		})
	}

	postID, err := post.Create(global.DBEngine, &images)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to create post", err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"post_id": postID}, "发布成功")
}

func GetPosts(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		response.Error(c, http.StatusOK, response.CodeInvalidParams, "参数错误", err.Error())
		return
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		response.Error(c, http.StatusOK, response.CodeInvalidParams, "参数错误", err.Error())
		return
	}

	p := model.Post{}
	posts, err := p.Get(global.DBEngine, pageInt, pageSizeInt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get posts", err.Error())
		return
	}

	var postList = []response.PostData{}
	for _, post := range *posts {
		// get images
		images, err := p.GetImages(global.DBEngine, post.ID)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get images", err.Error())
			return
		}

		var imgList []response.PostImage
		for _, img := range *images {
			imgList = append(imgList, response.PostImage{
				Url:    img.Url,
				Width:  img.Width,
				Height: img.Height,
				Index:  img.Index,
			})
		}

		// get user info
		u := model.User{
			Model: &gorm.Model{ID: post.UserID},
		}
		user, err := u.Get(global.DBEngine)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get user", err.Error())
			return
		}

		postList = append(postList, response.PostData{
			ID:          post.ID,
			UserID:      post.UserID,
			Avatar:      user.Avatar,
			Name:        user.Name,
			PublishedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
			Images:      imgList,
			Title:       post.Title,
			Content:     post.Content, Camera: func() string {
				if post.Camera == "" {
					return "N/A"
				} else {
					return post.Camera
				}
			}(),
			Lens: func() string {
				if post.Lens == "" {
					return "N/A"
				} else {
					return post.Lens
				}
			}(),
		})
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"total": len(postList), "posts": postList}, "获取成功")
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}

	p := model.Post{
		Model: &gorm.Model{ID: uint(idInt)},
	}

	post, err := p.GetByID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get post", err.Error())
		return
	}

	images, err := p.GetImages(global.DBEngine, post.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get images", err.Error())
		return
	}

	var imgList []response.PostImage
	for _, img := range *images {
		imgList = append(imgList, response.PostImage{
			Url:    img.Url,
			Width:  img.Width,
			Height: img.Height,
			Index:  img.Index,
		})
	}

	// get user info
	u := model.User{
		Model: &gorm.Model{ID: post.UserID},
	}
	user, err := u.Get(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get user", err.Error())
		return
	}

	postData := response.PostData{
		ID:          post.ID,
		UserID:      post.UserID,
		Avatar:      user.Avatar,
		Name:        user.Name,
		PublishedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		Images:      imgList,
		Title:       post.Title,
		Content:     post.Content, Camera: func() string {
			if post.Camera == "" {
				return "N/A"
			} else {
				return post.Camera
			}
		}(),
		Lens: func() string {
			if post.Lens == "" {
				return "N/A"
			} else {
				return post.Lens
			}
		}(),
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"post": postData}, "获取成功")
}

func DeletePost(c *gin.Context) {
	userData, _ := c.Get("user")
	role := userData.(map[string]string)["role"]
	uid := userData.(map[string]string)["id"]
	uidInt, _ := strconv.Atoi(uid)

	postId := c.Param("id")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}

	p := model.Post{
		Model: &gorm.Model{ID: uint(postIdInt)},
	}
	post, err := p.GetByID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get post", err.Error())
		return
	}

	// check permission
	if role != common.Admin.String() && post.UserID != uint(uidInt) {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "permission denied", "")
		return
	}

	if err := p.Delete(global.DBEngine); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to delete post", err.Error())
		return
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{}, "删除成功")
}

func GetUserPosts(c *gin.Context) {
	uid := c.Param("id")
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidParams, "invalid id", err.Error())
		return
	}

	p := model.Post{
		UserID: uint(uidInt),
	}
	posts, err := p.GetByUserID(global.DBEngine)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get posts", err.Error())
		return
	}

	var postList = []response.PostData{}
	for _, post := range *posts {
		// get images
		images, err := p.GetImages(global.DBEngine, post.ID)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get images", err.Error())
			return
		}

		var imgList []response.PostImage
		for _, img := range *images {
			imgList = append(imgList, response.PostImage{
				Url:    img.Url,
				Width:  img.Width,
				Height: img.Height,
				Index:  img.Index,
			})
		}

		// get user info
		u := model.User{
			Model: &gorm.Model{ID: post.UserID},
		}
		user, err := u.Get(global.DBEngine)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, response.CodeServerBusy, "failed to get user", err.Error())
			return
		}

		postList = append(postList, response.PostData{
			ID:          post.ID,
			UserID:      post.UserID,
			Avatar:      user.Avatar,
			Name:        user.Name,
			PublishedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
			Images:      imgList,
			Title:       post.Title,
			Content:     post.Content,
			Camera: func() string {
				if post.Camera == "" {
					return "N/A"
				} else {
					return post.Camera
				}
			}(),
			Lens: func() string {
				if post.Lens == "" {
					return "N/A"
				} else {
					return post.Lens
				}
			}(),
		})
	}

	response.Success(c, http.StatusOK, response.CodeSuccess, response.Data{"total": len(postList), "posts": postList}, "获取成功")
}
