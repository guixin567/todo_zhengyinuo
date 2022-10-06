package mysql

import (
	"errors"
	"todo_zhengyinuo/domain"
)

func CommunityCategory() (communities []domain.Community, err error) {

	db.Find(&communities)
	if len(communities) == 0 {
		err = errors.New("community table is empty")
	}
	return
}
