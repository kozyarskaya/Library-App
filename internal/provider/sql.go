package provider

import (
	"HW/internal/api"
	"database/sql"
	"errors"
	"fmt"
)

func (p *Provider) SelectArticle(id int) (api.Article, error) {
	fmt.Println("in select art")
	var art api.Article
	err := p.articlesDB.QueryRow("SELECT id, title, content, author_id FROM articles WHERE id = $1 AND deleted = FALSE", id).Scan(&art.Id, &art.Title, &art.Content, &art.AuthorId)
	if err != nil {
		if err == sql.ErrNoRows {
			return api.Article{}, errors.New("article not found") // Возвращаем ошибку, если статья не найдена
		}
		return api.Article{}, err // Возвращаем ошибку при других проблемах
	}
	fmt.Println(art.Content)
	return art, nil //
}

func (p *Provider) InsertArticle(a api.Article) error {
	_, err := p.articlesDB.Exec("INSERT INTO articles (title, content, author_id, deleted) VALUES ($1, $2, $3, FALSE)", a.Title, a.Content, a.AuthorId)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) UpdateArticle(a api.Article) error {
	_, err := p.articlesDB.Exec("UPDATE articles SET title = $1, content = $2, author_id = $3 WHERE id = $4 AND deleted = FALSE", a.Title, a.Content, a.AuthorId, a.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) DeleteArticle(id int) error {
	_, err := p.articlesDB.Exec("UPDATE articles SET deleted = TRUE WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) CheckUser(u api.User) (api.User, error) {
	var user api.User
	err := p.usersDB.QueryRow(`SELECT id, email, name FROM users WHERE email = $1`, u.Email).Scan(&user.Id, &user.Email, &user.Name)

	fmt.Println("check user")

	if err != nil {
		if err == sql.ErrNoRows {
			// Если пользователь не найден, возвращаем пустого пользователя и nil как ошибку
			return api.User{}, nil
		}
		// Возвращаем ошибку при проблеме с запросом
		return api.User{}, fmt.Errorf("ошибка при проверке пользователя: %w", err)
	}
	fmt.Println("checked from check")
	// Если пользователь найден, возвращаем его и ошибку о том, что почта занята
	return user, nil // Возвращаем найденного пользователя без ошибки
}

// Определение ошибки для занятости email
var ErrEmailAlreadyTaken = errors.New("email уже занят")

func (p *Provider) CreateUser(u api.User) error {
	_, err := p.usersDB.Exec("INSERT INTO users (name, email, hashedPassword) VALUES ($1, $2, $3)", u.Name, u.Email, u.HashedPassword)
	fmt.Println("create user")
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) SelectUser(emai string) (api.User, error) {
	var user api.User
	err := p.usersDB.QueryRow("SELECT id, name, email, hashedPassword FROM users WHERE email = $1", emai).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return api.User{}, errors.New("invalid credentials") // Пользователь не найден
		}
		return api.User{}, err // Возвращаем ошибку при проблеме с запросом
	}
	return user, nil
}
