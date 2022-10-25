package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)



//创建post
func CreatePostHandler(c *gin.Context) {
	//获取参数及校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("ShouldBindJSON", zap.Any("err", err))
		zap.L().Error("create post invalid  failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	//从 c 取到当前发请求的用户的ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorId = userID
	//创建帖子
	if err := logic.CreatPost(p); err != nil {
		zap.L().Error("logic.CreatPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

//post 详情
func PostDetaliHandler(c *gin.Context) {
	//获取参数（从url中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	//根据id取出帖子数据
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

//获取帖子列表
func GetPostListHandler(c *gin.Context) {
	//获取分页数据
	page, size := getPageInfo(c)
	//获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// 根据前端传来的参数动态的获取帖子列表
// 按照时间排序或者按照分数排序
/*
1.获取参数
2.去redis查询id列表
3.根据id去数据库查询帖子详细信息
*/

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	// ShouldBindQuery 是根据传过来的 query string 找到对应的参数
	// ShouldBindJSON  是根据传过来的 json找到对应的参数
	// ShouldBind  根据请求的数据类型动态的获取数据
	if err := c.ShouldBindQuery(p);err!= nil{
		zap.L().Error("GetPostListHandler2 invalid param",zap.Error(err))
		ResponseError(c,CodeInvalidParams)
		return
	}
	//获取分页数据
	//page, size := getPageInfo(c)
	//获取数据
	data, err := logic.GetPostListNew(p)   //更新：合二为一
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// 根据社区去查询帖子列表
//func GetCommunityPostListHandler(c *gin.Context) {
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	// ShouldBindQuery 是根据传过来的 query string 找到对应的参数
//	// ShouldBindJSON  是根据传过来的 json找到对应的参数
//	// ShouldBind  根据请求的数据类型动态的获取数据
//	if err := c.ShouldBindQuery(p);err!= nil{
//		zap.L().Error("GetCommunityPostListHandler invalid param",zap.Error(err))
//		ResponseError(c,CodeInvalidParams)
//		return
//	}
//	//获取分页数据
//	//page, size := getPageInfo(c)
//	//获取数据
//	data, err := logic.GetCommunityPostList2(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	//返回响应
//	ResponseSuccess(c, data)
//}
