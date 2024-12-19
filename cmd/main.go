package main

import (
	"HW/internal/api"
	"HW/internal/config"
	"HW/internal/provider"
	"HW/internal/usecase"
	"flag"
	"log"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	//"html/template"
)

func main() {

	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "D:\\Go\\go1.21.1\\src\\HW\\\\configs\\config.yaml", "путь к файлу конфигурации")
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
		AllowOrigins:  []string{"*"}, // Замените * на конкретные домены для большей безопасности
		AllowMethods:  []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        300,
	}))
	srv.Run()
}
