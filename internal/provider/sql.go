package provider

import (
	"HW/internal/api"
	"database/sql"
	"errors"
	"fmt"
)

func (p *Provider) SelectArticle(id int) (api.ArticleRes, error) {
	fmt.Println("in select art")
	var art api.ArticleRes
	au_id := 0
	err := p.articlesDB.QueryRow("SELECT id, title, text, author_id, date FROM articles  WHERE id = $1 AND deleted = FALSE", id).Scan(&art.Id, &art.Title, &art.Text, &au_id, &art.Date)
	err1 := p.usersDB.QueryRow("SELECT name FROM users  WHERE id = $1", au_id).Scan(&art.AuthorName)
	fmt.Println("scanned", art)
	if err != nil || err1 != nil {
		if err == sql.ErrNoRows {
			return api.ArticleRes{}, errors.New("article not found") // Возвращаем ошибку, если статья не найдена
		}
		return api.ArticleRes{}, err // Возвращаем ошибку при других проблемах
	}
	fmt.Println("art.Text")
	return art, nil //
}

func (p *Provider) SelectFullArticle(id int) (api.Article, error) {
	fmt.Println("in select art")
	var art api.Article
	err := p.articlesDB.QueryRow("SELECT id, title, content, author_name,author_id, date FROM articles INNER JOIN users ON users.id = articles.author_id  WHERE id = $1 AND deleted = FALSE", id).Scan(&art.Id, &art.Title, &art.Text, &art.AuthorName, &art.AuthorId, &art.Date)
	if err != nil {
		fmt.Println("not found")
		if err == sql.ErrNoRows {
			return api.Article{}, errors.New("article not found") // Возвращаем ошибку, если статья не найдена
		}
		return api.Article{}, err // Возвращаем ошибку при других проблемах
	}
	fmt.Println(art.Text)
	return art, nil //
}

func (p *Provider) SelectIds() ([]int, error) {
	var ids []int
	rows, err := p.articlesDB.Query("SELECT id FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем rows после завершения работы с ними
	fmt.Println("SelectIds")
	// Проходим по результатам запроса и добавляем ID в массив
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err // Возвращаем ошибку при сканировании строки
		}
		ids = append(ids, id) // Добавляем ID в массив
	}
	fmt.Println(ids)
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("1")
	return ids, nil
}

// ИСПРАВИТЬ!
func (p *Provider) InsertArticle(a api.Article) error {
	fmt.Println("InsertArticle")
	_, err := p.articlesDB.Exec("INSERT INTO articles (title, text, author_id, date, deleted) VALUES ($1, $2, $3, $4, FALSE)", a.Title, a.Text, a.AuthorId, a.Date)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// ИСПРАВИТЬ!
func (p *Provider) UpdateArticle(a api.Article) error {
	_, err := p.articlesDB.Exec("UPDATE articles SET title = $1, content = $2, author_id = $3 WHERE id = $4 AND deleted = FALSE", a.Title, a.Text, a.AuthorId, a.Id)
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

//func (p *Provider) CheckIsAuthor() (bool, error)
