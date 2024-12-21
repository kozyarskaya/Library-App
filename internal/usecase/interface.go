package usecase

import "HW/internal/api"

type Provider interface {
	SelectArticle(int) (api.ArticleRes, error)
	SelectFullArticle(int) (api.Article, error)
	InsertArticle(api.Article) error
	UpdateArticle(api.Article) error
	DeleteArticle(int) error
	SelectIds() ([]int, error)
	CheckUser(api.User) (api.User, error)
	CreateUser(api.User) error
	SelectUser(string) (api.User, error)
}
