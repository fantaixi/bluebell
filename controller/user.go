package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//处理注册请求
func SignUpHandler(c *gin.Context) {
	// 1、获取参数和参数校验
	 p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p);err != nil {
		//请求参数有误，直接返回
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		//判断err是不是validator.ValidationErrors的类型
		errs,ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK,gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)),  //对英文错误进行翻译
		//})
		return
	}
	//手动对请求参数进行详细的业务规则校验
	//用validator代替，也就是tag里面的binding
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(p)
	// 2、业务处理
	if err := logic.SignUp(p);err != nil{
		zap.L().Error("logic.SignUp failed",zap.Error(err))
		if  errors.Is(err,mysql.ErrorUserExit){
			ResponseError(c,CodeUserExist)
		}
		ResponseError(c,CodeServerBusy)
		return
	}
	// 3、返回响应
	//c.JSON(http.StatusOK,gin.H{
	//	"msg": "success",
	//})
	ResponseSuccess(c,nil)
}

//处理登录请求
func LoginHandler(c *gin.Context) {
	// 1、获取参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p);err != nil{
		//请求参数有误，直接返回
		zap.L().Error("Login with invalid param",zap.Error(err))
		//判断err是不是validator.ValidationErrors的类型
		errs,ok := err.(validator.ValidationErrors)
		if !ok {
			//参数错误
			ResponseError(c,CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK,gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)),  //对英文错误进行翻译
		//})
		return
	}
	// 2、业务处理
	user, err := logic.Login(p)
	if err!= nil{
		zap.L().Error("logic.Login failed",zap.String("username",p.Username),zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExit) {
			ResponseError(c,CodeUserNotExist)
			return
		}
		ResponseError(c,CodeInvalidPassword)
		//c.JSON(http.StatusOK,gin.H{
		//	"msg":"用户名或密码错误",
		//})
		return
	}
	// 3、返回响应
	//c.JSON(http.StatusOK,gin.H{
	//	"msg":"登录成功",
	//})
	ResponseSuccess(c,gin.H{
		"user_id": fmt.Sprintf("%d",user.UserID),  //id值会大于 1 << 53-1
		"user_name": user.Username,
		"token": user.Token,
	})
}