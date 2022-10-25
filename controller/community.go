package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//  社区相关
func CommunityHandler(c *gin.Context) {
	//查询到所有的社区(community_id,community_name) 以列表的形式返回
	data,err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)  //不把服务器错误返回给前端
		return
	}
	ResponseSuccess(c,data)
}

//  根据ID拿到社区详情
func CommunityDetailHandler(c *gin.Context) {
	//获取ID
	idStr := c.Param("id") //获取url参数
	id,err := strconv.ParseInt(idStr,10,64)
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	//根据id获取详情
	data,err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)  //不把服务器错误返回给前端
		return
	}
	ResponseSuccess(c,data)
}
