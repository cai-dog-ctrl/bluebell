package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) error {
	post.PostID = snowflake.GenID()
	return mysql.CreatePost(post)
}

func GetPostByID(ID int64) (data *models.ApiPost, err error) {
	Post, err := mysql.GetPostById(ID)
	if err != nil {
		zap.L().Error("mysql.GetPostById", zap.Error(err))
		return nil, err
	}
	username, err := mysql.GetUserByID(Post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID", zap.Error(err))
		return nil, err
	}
	Community, err := mysql.GetCommunityDetailByID(Post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID", zap.Error(err))
		return nil, err
	}
	data = new(models.ApiPost)
	data.Post = Post
	data.AuthorName = username
	data.CommunityDetail = Community
	return
}

func GetPostList(pageSize, pageNum int64) (data []*models.ApiPost, err error) {
	Post, err := mysql.GetPostList(pageSize, pageNum)
	if err != nil {
		zap.L().Error("mysql.GetPostList", zap.Error(err))
		return nil, err
	}
	data = make([]*models.ApiPost, 0, len(Post))
	for _, v := range Post {
		username, err := mysql.GetUserByID(v.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID", zap.Error(err))
			return nil, err
		}
		Community, err := mysql.GetCommunityDetailByID(v.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID", zap.Error(err))
			return nil, err
		}
		postDetail := &models.ApiPost{
			AuthorName:      username,
			Post:            v,
			CommunityDetail: Community,
		}
		data = append(data, postDetail)
	}
	return
}
