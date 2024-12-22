package main

import (
	"Library-App-dev_new/internal/api"
	"Library-App-dev_new/internal/config"
	"Library-App-dev_new/internal/provider"
	"Library-App-dev_new/internal/usecase"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	//"html/template"
)

func main() {

	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "D:\\Progects Go\\Library-App-dev_new\\\\configs\\config.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	//Инициализация провайдера
	usersProvider := provider.NewProvider(cfg.DB.Users, cfg.DB.Articles)
	// Инициализация бизнес-логики с использованием обоих провайдеров
	use := usecase.NewUsecase(usersProvider)

	srv := api.NewServer(cfg.IP, cfg.Port, use)
	srv.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                                                // Разрешить все источники (можно указать конкретные)
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, // Разрешенные методы
		AllowHeaders: []string{"Content-Type", "Authorization"},                                    // Разрешенные заголовки
	}))

	fmt.Println("Stert!")
	srv.Run()
}
