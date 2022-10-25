package models

//定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

//注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//投票
type ParamVoteData struct {
	//user_id 可以根据传进来的参数获取，所以不需要再定义一次
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
}

// 帖子列表 query string 参数
type ParamPostList struct {
	//Page  int64  `json:"page" form:"page"`
	//Size  int64  `json:"size" form:"size"`
	//Order string `json:"order" form:"order"`
	//CommunityID int64 `json:"community_id" form:"community_id"` //可以为空

	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

// 帖子列表 query string 参数
//type ParamCommunityPostList struct {
//	*ParamPostList
//
//}
