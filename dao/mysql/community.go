package mysql

import (
	"todo_zhengyinuo/domain"
)

func CommunityCategory() (communities []domain.Community, err error) {

	find := db.Find(&communities)
	err = find.Error
	return
}

func CommunityCategoryDetail(communityId int64) (*domain.Community, error) {
	community := &domain.Community{CommunityId: communityId}
	first := db.Where(community).First(&community)
	return community, first.Error
}
