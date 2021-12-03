package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func SingUpHandler(c *gin.Context) {
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp With invalid param", zap.Error(err))
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}
	if len(p.Username) == 0 || len(p.RePassword) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("SignUp With invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.Signup failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamsLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login With invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名密码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
