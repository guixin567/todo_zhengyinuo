package domain

type RegisterParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddArticleParam struct {
	AuthorId      int64  `json:"author_id"`
	ArticleId     int64  `json:"article_id" binding:"required"`
	CommunityId   int64  `json:"community_id" binding:"required"`
	CommunityName string `json:"community_name" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Content       string `json:"content" binding:"required"`
}
