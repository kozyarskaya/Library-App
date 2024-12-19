package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Server  *echo.Echo
	address string
	uc      Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	api := Server{
		uc: uc,
	}
	api.Server = echo.New()
	api.Server.POST("/registration", api.signUp)
	api.Server.POST("/auth/sign-in", api.signIn)

	api.Server.POST("/api/articles/", api.postArticle)

	api.Server.GET("/api/articles/:id", api.getArticleById)
	api.Server.PUT("/api/articles/:id", api.putArticle)
	api.Server.DELETE("/api/articles/:id", api.deleteArticle)

	api.address = fmt.Sprintf("%s:%d", ip, port)
	fmt.Println("Routes registered:")
	for _, route := range api.Server.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
	return &api
}

func (api *Server) Run() {
	api.Server.Logger.Fatal(api.Server.Start(api.address))
}
