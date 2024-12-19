package usecase

// реализуем бизнес-логику, ей вызывают обрабтчики запросов, а она передает задачу бд
import (
	"HW/internal/api"
	"fmt"
)

func (u *Usecase) CreateArticle(v api.Article) error {
	err := u.p.UpdateArticle(v)
	if err != nil {
		return err // Возвращаем пустую статью и ошибку
	}
	return nil // Возвращаем созданную статью и nil для ошибки
}

func (u *Usecase) FetchArticleByID(id int) (api.Article, error) {
	msg, err := u.p.SelectArticle(id)
	if err != nil {
		fmt.Println(id, err)
		return api.Article{}, err
	}
	fmt.Println(id)
	return msg, nil
}

func (u *Usecase) ChangeArticleByID(v api.Article) error {
	err := u.p.InsertArticle(v)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) DeleteArticleByID(id int) error {
	err := u.p.DeleteArticle(id)
	if err != nil {
		return err
	}
	return nil
}
