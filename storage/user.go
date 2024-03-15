package storage

import (
	"context"

	"example.com/gin_forum/models"
)

/*
*
create table user
(

	username varchar(128)             not null
	    primary key,
	password varchar(512)  default '' not null,
	email    varchar(256)  default '' not null,
	image    varchar(1024) default '' not null,
	bio      varchar(1024) default '' not null,
	constraint email
	    unique (email)

)

	charset = utf8mb4;
*/
func CreateUser(ctx context.Context, user *models.User) error {
	_, err := db.ExecContext(ctx, "insert user(username, password, email, image, bio) values(?, ?, ?, ?, ?)",
		user.Username, user.Password, user.Email, user.Image, user.Bio)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := db.GetContext(ctx, &user, "select username, password, email, image, bio from user where email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := db.GetContext(ctx, &user, "select username, password, email, image, bio from user where username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := db.ExecContext(ctx, "delete from user where email = ?", email)
	return err
}

func UpdateUserByUsername(ctx context.Context, username string, user *models.User) error {
	if user.Password != "" {
		_, err := db.ExecContext(ctx, "update user set username=?, password=?, email=?, image=?, bio=? where username=?",
			user.Username, user.Password, user.Email, user.Image, user.Bio, username)
		return err
	}
	_, err := db.ExecContext(ctx, "update user set username=?, email=?, image=?, bio=? where username=?",
		user.Username, user.Email, user.Image, user.Bio, username)
	return err
}
