package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated = errors.New("不允许重复投票")
)

func CreatePost(postID,communityID int64) (err error){
	piepline := client.TxPipeline()
	// 帖子时间
	 piepline.ZAdd(getRedisKey(KeyPostTimeZSet),redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	 piepline.ZAdd(getRedisKey(keyPostScoreZSet),redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(communityID)))
	piepline.SAdd(cKey,postID)
	_, err = piepline.Exec()
	return
}

func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票限制
	//去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2.更新帖子的分数
	//先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	//如果这一次的投票的值和之前的保存的一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值
	// 2和3 需要一个piepeline事务操作
	piepeline := client.TxPipeline()
	piepeline.ZIncrBy(getRedisKey(keyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3.记录用户为该帖子投票的数据
	if value == 0 {
		piepeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), postID)
	} else {
		piepeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, //赞成票还是反对票
			Member: userID,
		})
	}
	_, err := piepeline.Exec()
	return err
}
