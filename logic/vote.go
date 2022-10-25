package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

//为帖子投票
/*
投票的几种情况
direction = 1
	之前没有投票，现在投赞成票
	之前投反对票，现在投赞成票
= 0
	之前投过赞成票，现在要取消投票
	之前投过反对票，现在要取消投票
=-1
	之前没有投票，现在投反对票
	之前投赞成票，现在投反对票

投票限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票
	1.到期之后就将redis中保存的赞成票及反对票存到mysql中
	2.到期之后删除 KeyPostVotedZSetPF
 */
func VoteForPost(userID int64,p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID",userID),
		zap.String("postID",p.PostID),
		zap.Int8("Direction",p.Direction),
		)
	return redis.VoteForPost(strconv.Itoa(int(userID)),p.PostID,float64(p.Direction))
}
