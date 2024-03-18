package models

import "time"

/**
create table article_comment
(
    id              bigint auto_increment
        primary key,
    author_username varchar(255)                            not null,
    body            text                                    not null,
    created_at      datetime      default CURRENT_TIMESTAMP not null,
    updated_at      datetime      default CURRENT_TIMESTAMP not null
)
    charset = utf8mb4;
*/

type ArticleComment struct {
	Id             int64  `gorm:"primaryKey"`
	AuthorUsername string `gorm:"column:author_username"`
	Body           string
	ArticleId      int64
	CreatedAt      time.Time
	UpdatedAt      time.Time

	AuthorUserEmail string `gorm:"->"`
	AuthorUserImage string `gorm:"->"`
	AuthorUserBio   string `gorm:"->"`
}

func (ac ArticleComment) TableName() string {
	return "article_comment"
}
