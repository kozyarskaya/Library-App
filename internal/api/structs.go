package api

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id             int    `json:"-"`
	Name           string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
}

//LoginRequest
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}
type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Article struct {
	Id       int    `json:"-"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int    `json:"authorId"`
	Deleted  bool   `json:"deleted"`
}

type Response struct {
	Ok      string `json:"ok"`
	Message string `json:"message"`
}
