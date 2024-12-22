package api

type Usecase interface {
	FetchArticleByID(int) (ArticleRes, error)
	FetchFullArticleByID(int) (Article, error)
	ChangeArticleByID(Article) error
	DeleteArticleByID(int) error
	CreateArticle(Article) error
	FetchAllId() ([]int, error)
	SignUp(User) (string, error)
	SignIn(Credentials) (string, error) //Возвращает JWT токен,
}

type Authorization interface {
}
