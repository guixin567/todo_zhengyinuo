package domain

type Article struct {
	ID            int64  `gorm:"column:id" db:"id" json:"id"`
	CommunityId   int64  `gorm:"column:community_id" db:"community_id" json:"community_id"`
	CommunityName string `gorm:"column:community_name" db:"community_name" json:"community_name"`
	AuthorId      int64  `gorm:"column:author_id" db:"author_id" json:"author_id"`
	ArticleId     int64  `gorm:"column:article_id" db:"article_id" json:"article_id"`
	Title         string `gorm:"column:title" db:"title" json:"title"`
	Content       string `gorm:"column:content" db:"content" json:"content"`
	Status        int64  `gorm:"column:status" db:"status" json:"status"`
}

func (Article) TableName() string {
	return "article"
}
