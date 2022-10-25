package redis

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"   //zset; 帖子及发帖时间
	keyPostScoreZSet   = "post:score"  //zset;帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" //zset;记录用户及投票的类型,参数是post_id

	KeyCommunitySetPF = "community:"  //set; 保存每个分区下帖子的id
)

// 给 redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix+key
}
