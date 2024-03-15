package storage

import (
	"context"
	"testing"

	"example.com/gin_forum/models"
	"example.com/gin_forum/utils"
)

func TestCreateAndGetUser(t *testing.T) {
	ctx := context.Background()
	userName := "xxx33x123"
	email := "xxxx33123@gmail.com"

	err := CreateUser(ctx, &models.User{
		Username: userName,
		Password: "xxxx123",
		Email:    email,
		Image:    "xxxx123",
		Bio:      "xxxx123",
	})
	if err != nil {
		t.Errorf("create user failed, err: %v", err)
		return
	}

	dbUser, err := GetUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("get user by email failed, err: %v", err)
		return
	}

	t.Logf("get user: %v\n", utils.JsonMarshal(dbUser))

	err = DeleteUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("delete user by email failed")
		return
	}

}
