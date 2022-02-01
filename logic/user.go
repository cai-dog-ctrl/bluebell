package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p models.ParamSignUp) (err error) {
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	//生成UID
	userID := snowflake.GenID()
	u := models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}

	return mysql.InsertUser(&u)
}
func Login(p *models.ParamsLogin) (user *models.User, err error) {
	user = new(models.User)
	user.UserName = p.Username
	user.Password = p.Password
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}
