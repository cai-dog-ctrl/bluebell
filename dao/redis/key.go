package redis

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"  //帖子及发帖时间
	KeyPostScoreZSet   = "post:score" //帖子投票分数
	KeyPostVotedZSetPF = "post:voted" //记录用户及投票类型 ，参数是post ID
)

func GetRedisKey(Key string) string {
	return KeyPrefix + Key
}
