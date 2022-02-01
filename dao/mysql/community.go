package mysql

import (
	"bluebell/models"
	"database/sql"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlstr := "select community_id,community_name from community"
	err = db.Select(&communityList, sqlstr)
	return
}
func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {
	community := new(models.CommunityDetail)
	sqlStr := "select community_id ,community_name,introduction,create_time from community where community_id= ?"
	err := db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return nil, err
	}
	return community, nil
}
