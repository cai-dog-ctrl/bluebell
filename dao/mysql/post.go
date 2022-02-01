package mysql

import "bluebell/models"

func CreatePost(post *models.Post) error {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id) VALUES (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, post.PostID, post.Title, post.Content, post.AuthorId, post.CommunityID)
	return err
}

func GetPostById(ID int64) (*models.Post, error) {
	data := &models.Post{}
	sqlStr := "select post_id,title,content,author_id,community_id,create_time from post where post_id = ?"
	err := db.Get(data, sqlStr, ID)
	return data, err
}

func GetUserByID(Id int64) (username string, err error) {

	sqlStr := "select username from user where user_id = ?"
	err = db.Get(&username, sqlStr, Id)
	return
}

func GetPostList(pageSize, pageNum int64) (data []*models.Post, err error) {
	data = make([]*models.Post, 0, 2)
	sqlStr := " select post_id,title,content,author_id,community_id,create_time from post limit ?,?"
	err = db.Select(&data, sqlStr, (pageNum-1)*pageSize, pageSize)
	return
}
