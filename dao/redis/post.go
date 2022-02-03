package redis

import (
	"bluebell/models"
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	OneWeekForSecond = 7 * 24 * 3600
)

var (
	ErrPostTimeExpire = errors.New("帖子投票时间已过")
)

func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(GetRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	pipeline.ZAdd(GetRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	_, err := pipeline.Exec()
	return err

}
func VoteForPost(UserID, PostID string, v float64) error {
	//获得帖子发帖时间
	postTime := rdb.ZScore(GetRedisKey(KeyPostTimeZSet), PostID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekForSecond {
		return ErrPostTimeExpire
	}

	ov := rdb.ZScore(GetRedisKey(KeyPostScoreZSet+PostID), UserID).Val()
	var dir float64
	if v > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - v)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(GetRedisKey(KeyPostScoreZSet), dir*diff*432, PostID)

	if v == 0 {
		pipeline.ZRem(GetRedisKey(KeyPostVotedZSetPF+PostID), PostID)

	} else {
		pipeline.ZAdd(GetRedisKey(KeyPostVotedZSetPF+PostID), redis.Z{
			Score:  v,
			Member: UserID,
		})
	}
	_, err := pipeline.Exec()
	return err

}

func GetPostIDsInOrder(p *models.ParmaPostList) ([]string, error) {
	key := GetRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = GetRedisKey(KeyPostTimeZSet)
	}
	start := (p.PageNum - 1) * p.PageSize
	end := start + p.PageSize - 1
	return rdb.ZRevRange(key, start, end).Result()

}
