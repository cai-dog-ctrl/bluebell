package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamsVoteData struct {
	PostID    string  `json:"post_id" binding:"required"`
	Direction float64 `json:"direction" binding:""`
}

type ParmaPostList struct {
	PageNum  int64  `json:"page_num" form:"page"`
	PageSize int64  `json:"page_size" form:"size"`
	Order    string `json:"order" form:"order"`
}
