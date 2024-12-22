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

//middleware - функция, которая при обработке запросов может выполнять доп действия до или после вызова обработчика

// проверяет наличие и валидность JWT в заголовке Authorization каждого запроса
// next - следдующий обработчик
// возвращает новую функцию, соответствующую типу echo.HandlerFunc, которая будет использоваться как middleware

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
	token, err := srv.uc.SignUp(user)
	if err != nil {
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
		Message: token,
	})
}

// Обработчик для аутентификации пользователя
func (srv *Server) signIn(c echo.Context) error {
	var credentials Credentials
	if err := c.Bind(&credentials); err != nil {
		fmt.Println("Invalid input from In")
		return c.JSON(http.StatusBadRequest, Response{
			Ok:      "false",
			Message: " Ошибка передачи параметров",
		})
	}
	fmt.Println("user signIn from api")
	maybe_token_msg, err := srv.uc.SignIn(credentials)
	if err != nil {
		fmt.Println("Authentication failed", err.Error())
		return c.JSON(http.StatusUnauthorized, Response{
			Ok:      "false",
			Message: maybe_token_msg,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Ok:      "true",
		Message: maybe_token_msg,
	})
}

func (srv *Server) postArticle(e echo.Context) error {
	fmt.Println("postArticle")
	userId := ""
	if authorId, ok := e.Get("authorId").(string); ok {
		userId = authorId
	} else {
		fmt.Println(authorId)
		// Обработка случая, когда authorId отсутствует или не является строкой
		return e.JSON(http.StatusUnauthorized, Response{
			Ok:      "false",
			Message: "Необходима авторизация",
		})
	}
	fmt.Println("qqq1")

	// Создаём новую статью для связывания данных из запроса
	var newArticle Article
	// Извлекаем данные из тела запроса и связываем их со структурой
	if err := e.Bind(&newArticle); err != nil {
		fmt.Println("qqq")

		return e.JSON(http.StatusBadRequest, Response{
			Ok:      "false",
			Message: "ошибка считывания",
		})
	}
	fmt.Println("qqq2")
	newArticle.AuthorId, _ = strconv.Atoi(userId)
	// Создаем статью в бизнес-логике
	err := srv.uc.CreateArticle(newArticle)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Response{
			Ok:      "false",
			Message: "ошибка создания статьи",
		})
	}
	return e.JSON(http.StatusCreated, Response{
		Ok:      "true",
		Message: "Статья создана!",
	})
}

// вернуть список json
func (srv *Server) getArticleById(e echo.Context) error {
	fmt.Println("getArticleById")
	ids, err := srv.uc.FetchAllId()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Response{
			Ok:      "false",
			Message: "Ошибка при получении ID статей",
		})
	}
	var ans []ArticleRes
	for _, id := range ids {
		article, err := srv.uc.FetchArticleByID(id)
		if err != nil {
			fmt.Printf("Статья с ID %d не найдена: %v\n", id, err)
			continue
		}
		ans = append(ans, article)
	}
	return e.JSON(http.StatusOK, ResponseN{
		Ok:      "true",
		Message: ans,
	})
}

func (srv *Server) getONEArticleById(e echo.Context) error {
	fmt.Println("getONEArticleById")
	// Извлекаем ID статьи из параметров запроса
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Ok:      "false",
			Message: "некорректный id статьи",
		})
	}
	// Получаем статью из бизнес-логики
	article, err := srv.uc.FetchArticleByID(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, Response{
			Ok:      "false",
			Message: "Статья не найдена",
		})
	}
	// Возвращаем статью в формате JSON
	return e.JSON(http.StatusOK, ResponseN{
		Ok:      "true",
		Message: article,
	})
}

//проверка доступа!!!!

func (srv *Server) putArticle(e echo.Context) error {
	userIdStr := e.Get("userId").(string)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid user ID"})
	}

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
	existingArticle, err := srv.uc.FetchFullArticleByID(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{"error": "Article not found"})
	}

	// Проверяем, является ли текущий пользователь автором статьи
	if existingArticle.AuthorId != userId {
		return e.JSON(http.StatusForbidden, echo.Map{"error": "not authorized to edit this article"})
	}

	// Устанавливаем ID обновляемой статьи
	updatedArticle.Id = id

	// Обновляем статью в бизнес-логике
	err = srv.uc.ChangeArticleByID(updatedArticle)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update article"})
	}

	// Возвращаем статус 204 No Content при успешном обновлении
	return e.NoContent(http.StatusOK)
}

//проверка доступа!!!!

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
