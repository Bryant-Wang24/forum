package storage

import (
	"context"

	"example.com/gin_forum/models"
	"example.com/gin_forum/params/request"
	"gorm.io/gorm"
)

func CreateArticle(ctx context.Context, article *models.Article) error {
	return gormDb.WithContext(ctx).Create(article).Error
}

func ListArticles(ctx context.Context, req *request.ListArticleQuery) ([]*models.Article, error) {
	var articles []*models.Article
	//gormDb.Raw()
	err := listArticleTx(ctx, req).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func listArticleTx(ctx context.Context, req *request.ListArticleQuery) *gorm.DB {
	tx := gormDb.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Order("created_at desc").
		Offset(req.Offset).Limit(req.Limit)
	if req.Tag != "" {
		// where ()
		tx = tx.Where("article.tag_list like ?", "%\""+req.Tag+"\"%")
	}
	return tx
}

func CountArticles(ctx context.Context, req *request.ListArticleQuery) (int64, error) {
	var count int64
	//gormDb.Table("article").Count(&count)
	if err := listArticleTx(ctx, req).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
