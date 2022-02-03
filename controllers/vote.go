package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteControllers(c *gin.Context) {
	p := new(models.ParamsVoteData)
	err := c.ShouldBindJSON(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		errDate := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParams, errDate)
		return
	}
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	err = logic.VoteForPost(userID, p)
	if err != nil {
		zap.L().Error("logic.VoteForPost ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}
