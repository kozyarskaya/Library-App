package repository

type Authorization interface {
}

type Article interface {
}

type Repository struct {
	Authorization
	Article
}

func NewRepository() *Repository {
	return &Repository{}
}
