package service

import (
	"todo_zhengyinuo/dao/mysql"
	"todo_zhengyinuo/domain"
)

func AddArticle(param *domain.AddArticleParam) (*domain.Article, error) {
	return mysql.InsertArticle(param)
}
