package models

import "time"

type Post struct {
	PostID      int64     `json:"post_id" db:"post_id" `
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ApiPost struct {
	AuthorName       string `json:"author_name"`
	*Post            `json:"*_post"`
	*CommunityDetail `json:"*_community_detail"`
}
