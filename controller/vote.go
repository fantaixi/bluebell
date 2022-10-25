package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//投票

func PostVoteController(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs,ok := err.(validator.ValidationErrors)  //类型断言
		if !ok {
			ResponseError(c,CodeInvalidParams)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))  //翻译并去掉错误提示中的结构体标识
		ResponseErrorWithMsg(c,CodeInvalidParams,errData)
		return
	}

	//获取当前请求的用户ID
	userID,err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c,CodeNeedLogin)
		return
	}

	//投票的业务逻辑
	if err := logic.VoteForPost(userID, p);err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}
