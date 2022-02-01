package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreatPost(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	UserID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorId = UserID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostById(c *gin.Context) {
	postId := c.Param("id")
	id, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		zap.L().Error("id param error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	data, err := logic.GetPostByID(id)
	if err != nil {
		zap.L().Error("Get Post error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func GetPostList(c *gin.Context) {
	pageSize, PageNum := GetPageSizeAndPageNum(c)
	data, err := logic.GetPostList(pageSize, PageNum)
	if err != nil {
		zap.L().Error("logic.GetPostList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
