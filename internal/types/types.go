package types

import (
	"time"
)

type User struct {
	Id           int
	Email        string
	PasswordHash string
}

type CreateUserData struct {
	Id       int
	Email    string
	Username string
	Password string
	Token    *string
	Expired  *time.Time
}

type GetUserData struct {
	Password string
	Email    string
}

type Err struct {
	StatusCode int
	StatusText string
}

type ErrText struct {
	Username string
	Email    string
	Pass1    string
	Pass2    string
}

type UserService interface {
	CreateUser(user *CreateUserData) error
	CheckUserExists(user *CreateUserData) (bool, ErrText)
	CheckLogin(user *GetUserData) (int, error)
	AddToken(userid int, cookie string) error
	RemoveToken(token string) error
	GetUserByToken(token string) (user *User, err error)
}

type CreatePost struct {
	AuthorId   int
	AuthorName string
	Title      string
	Content    string
	Categories []string
}

type UserRepo interface {
	CreateUserDB(user *User)
	GetUserEmailDB(user string) error
	CheckLoginDB(user *GetUserData) (int, error)
	AddTokenDB(userid int, cookieToken string) error
	RemoveTokenDB(token string) error
	GetUserByToken(token string) (user *User, err error)
}
