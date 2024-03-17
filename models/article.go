package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

/*
*
create table article
(

	id              bigint auto_increment
	    primary key,
	author_username varchar(255)                            not null,
	title           varchar(1024)                           not null,
	slug            varchar(512)                            not null,
	body            text                                    not null,
	description     varchar(2048) default ''                not null,
	tag_list        varchar(1024) default '[]'              not null,
	created_at      datetime      default CURRENT_TIMESTAMP not null,
	updated_at      datetime      default CURRENT_TIMESTAMP not null,
	constraint slug
	    unique (slug)

)

	charset = utf8mb4;
*/
type Article struct {
	Id             int64  `db:"id"`
	AuthorUsername string `gorm:"column:author_username"`
	Title          string
	Slug           string
	Body           string
	Description    string
	TagList        TagList `gorm:"type:string"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	AuthorUserEmail string `gorm:"->"`
	AuthorUserImage string `gorm:"->"`
	AuthorUserBio   string `gorm:"->"`
}

func (a Article) TableName() string {
	return "article"
}

type TagList []string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *TagList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, j)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j TagList) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}
