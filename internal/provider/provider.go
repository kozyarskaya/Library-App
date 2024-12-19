package provider

import (
	"HW/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

type Provider struct {
	usersDB    *sql.DB
	articlesDB *sql.DB
}

// NewProvider создает новый экземпляр Provider с подключением к базе данных.
func NewProvider(usersConfig config.DBConfig, articlesConfig config.DBConfig) *Provider {
	usersPsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		usersConfig.Host, usersConfig.Port, usersConfig.User, usersConfig.Password, usersConfig.DBname)

	// Создание соединения с сервером PostgreSQL для пользователей
	usersConn, err := sql.Open("postgres", usersPsqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения с базой данных пользователей
	if err = usersConn.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Users DB:", usersConfig.DBname)

	articlesPsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		articlesConfig.Host, articlesConfig.Port, articlesConfig.User, articlesConfig.Password, articlesConfig.DBname)

	// Создание соединения с сервером PostgreSQL для статей
	articlesConn, err := sql.Open("postgres", articlesPsqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения с базой данных статей
	if err = articlesConn.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Articles DB:", articlesConfig.DBname)

	return &Provider{
		usersDB:    usersConn,
		articlesDB: articlesConn,
	}
}
