package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

//把每一步数据库操作封装成函数，等待logic层（也就是service层）根据业务需求调用

const secret = "fantaixi.com"

// 指定用户名的用户是否存在
func CheckUserExit(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return  err
	}
	if count>0 {
		return ErrorUserExit
	}
	return
}

func InsertUser(user *models.User) (err error){
	//对密码加密
	 user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error){
	oPassword := user.Password //用户登录的密码
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(user,sqlStr,user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExit
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorPasswordWrong
	}
	return 
}

func GetUserById(uid int64)(user *models.User,err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	err = db.Get(user,sqlStr,uid)
	return
}
