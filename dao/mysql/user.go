package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "lihao"

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (err error) {

	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 插入一个User
func InsertUser(user *models.User) (err error) {
	//密码加密
	password := encryptPassword(user.Password)

	sqlStr := "insert into user(user_id,username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	return
}

func encryptPassword(Password string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(Password)))
}
func Login(user *models.User) error {
	oPassword := user.Password
	sqlStr := "select user_id,username,password from user where username= ?"
	err := db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return nil
}
