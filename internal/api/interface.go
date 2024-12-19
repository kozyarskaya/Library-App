package api

type Usecase interface {
	FetchArticleByID(int) (Article, error)
	ChangeArticleByID(Article) error
	DeleteArticleByID(int) error
	CreateArticle(Article) error
	SignUp(User) error
	SignIn(Credentials) (string, error) //Возвращает JWT токен,
}

type Authorization interface {
}

type Articles interface {
}
