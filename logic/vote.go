package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"
)

func VoteForPost(userID int64, p *models.ParamsVoteData) error {
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, p.Direction)
}
