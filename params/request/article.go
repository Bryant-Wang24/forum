package request

type ListArticleQuery struct {
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	Tag    string `form:"tag"`
}
