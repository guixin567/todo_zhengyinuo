package domain

type Community struct {
	ID            int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	CommunityId   int64  `gorm:"column:community_id" db:"community_id" json:"community_id" form:"community_id"`
	CommunityName string `gorm:"column:community_name" db:"community_name" json:"community_name" form:"community_name"`
	Introduction  string `gorm:"column:introduction" db:"introduction" json:"introduction" form:"introduction"`
}

func (Community) TableName() string {
	return "community"
}
