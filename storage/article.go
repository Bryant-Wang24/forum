package storage

import (
	"context"

	"example.com/gin_forum/models"
)

func CreateArticle(ctx context.Context, article *models.Article) error {
	return gormDb.WithContext(ctx).Create(article).Error
}

func ListArticles(ctx context.Context, limit, offset int) ([]*models.Article, error) {
	var articles []*models.Article
	err := gormDb.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Order("created_at desc").
		Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func CountArticles(ctx context.Context) (int64, error) {
	var count int64
	//gormDb.Table("article").Count(&count)
	if err := gormDb.WithContext(ctx).Model(models.Article{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
