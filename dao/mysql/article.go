package mysql

import "todo_zhengyinuo/domain"

func InsertArticle(param *domain.AddArticleParam) (*domain.Article, error) {
	article := &domain.Article{
		CommunityId:   param.CommunityId,
		CommunityName: param.CommunityName,
		AuthorId:      param.AuthorId,
		ArticleId:     param.ArticleId,
		Title:         param.Title,
		Content:       param.Content,
	}
	create := db.Create(article)
	return article, create.Error
}

func SearchArticle(articleId int64) (*domain.Article, error) {
	article := &domain.Article{ArticleId: articleId}
	tx := db.Where(article).First(article)
	return article, tx.Error
}
