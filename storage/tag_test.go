package storage

import (
	"context"
	"testing"

	"example.com/gin_forum/utils"
)

func TestListTags(t *testing.T) {
	ctx := context.TODO()
	res, err := ListPopularTags(ctx)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("tags: %v\n", utils.JsonMarshal(res))
}
