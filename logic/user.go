package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
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
func Login(p *models.ParamsLogin) error {
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)

}
