package api

// обработка http запросов и переотправление на бизнес-логику(работу с бд)
import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var ErrEmailAlreadyTaken = errors.New("email уже занят")

// signUp - обработчик для регистрации пользователя
func (srv *Server) signUp(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Ok:      "false",
			Message: "некорекктное считывание данных",
		})
	}
	fmt.Println("user signUp from api")
	// Вызываем бизнес-логику для регистрации пользователя
	if err := srv.uc.SignUp(user); err != nil {
		fmt.Println(err.Error(), errors.Is(err, ErrEmailAlreadyTaken))
		if errors.Is(err, ErrEmailAlreadyTaken) { // Проверяем конкретную ошибку
			fmt.Println("rrr")
			return c.JSON(http.StatusConflict, Response{
				Ok:      "false",
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, Response{
			Ok:      "false",
			Message: "ошибка сервера",
		})
	}
	fmt.Println("signUp sucs")
	return c.JSON(http.StatusCreated, Response{
		Ok:      "true",
		Message: "Успешная регистрация",
	})
}

// Обработчик для аутентификации пользователя
func (srv *Server) signIn(c echo.Context) error {
	var credentials Credentials
	if err := c.Bind(&credentials); err != nil {
		fmt.Println("Invalid input")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	token, err := srv.uc.SignIn(credentials)
	if err != nil {
		fmt.Println("Authentication failed")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (srv *Server) postArticle(e echo.Context) error {
	fmt.Println("Received request for sign up")

	// Создаём новую статью для связывания данных из запроса
	var newArticle Article
	// Извлекаем данные из тела запроса и связываем их со структурой
	if err := e.Bind(&newArticle); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	// Создаем статью в бизнес-логике
	err := srv.uc.CreateArticle(newArticle)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create article"})
	}
	// Возвращаем созданную статью в формате JSON с кодом 201 Created
	return e.NoContent(http.StatusCreated)
}

func (srv *Server) getArticleById(e echo.Context) error {
	// Извлекаем ID статьи из параметров запроса
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid article ID"})
	}
	// Получаем статью из бизнес-логики
	article, err := srv.uc.FetchArticleByID(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{"error": "Article not found"})
	}
	// Возвращаем статью в формате JSON
	return e.JSON(http.StatusOK, article)
}

func (srv *Server) putArticle(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid article ID"})
	}
	var updatedArticle Article
	//связывания данных из HTTP-запроса с переменной
	if err := e.Bind(&updatedArticle); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	// Устанавливаем ID обновляемой статьи
	updatedArticle.Id = id
	// Обновляем статью в бизнес-логике
	err = srv.uc.ChangeArticleByID(updatedArticle)
	if err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{"error": "Article not found"})
	}
	// Возвращаем обновлённую статью в формате JSON
	return e.NoContent(http.StatusOK)
}

func (srv *Server) deleteArticle(e echo.Context) error {
	// Извлекаем ID статьи из параметров запроса
	idStr := e.Param("id")
	// Преобразуем строку в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid article ID"})
	}
	// Удаляем статью в бизнес-логике
	err = srv.uc.DeleteArticleByID(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{"error": "Article not found"})
	}
	// Возвращаем статус 204 No Content при успешном удалении
	return e.NoContent(http.StatusNoContent)
}
