package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error){
	//判断用户是否存在
	if err = mysql.CheckUserExit(p.Username);err != nil {
		//数据库查询出错
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	// 密码加密
	//构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//存入数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User,err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针，能拿到user.UserID
	if err:=mysql.Login(user);err!=nil{
		return nil,err
	}
	//生成token
	token,err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
