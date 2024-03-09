package user

import (
	"database/sql"
	"fmt"

	t "forum/internal/types"

	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	DB *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}

func (db *UserDB) CreateUserDB(user *t.User) {
	_, err := db.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)",
		user.Email,
		user.PasswordHash)

	fmt.Println(user.Email, user.PasswordHash)
	if err != nil {
		fmt.Println("repository user err:", err)
		return
	}
}

func (db *UserDB) CheckLoginDB(user *t.GetUserData) (int, error) {
	var userMatch t.CreateUserData
	err := db.DB.QueryRow("SELECT * FROM users WHERE email= $1", user.Email).Scan(
		&userMatch.Id,
		&userMatch.Email,
		&userMatch.Password,
		&userMatch.Token,
		&userMatch.Expired)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userMatch.Password), []byte(user.Password))

	if err != nil {
		return 0, err
	}
	return userMatch.Id, nil
}

func (db *UserDB) GetUserEmailDB(userEmail string) error {
	var userMatch t.CreateUserData
	err := db.DB.QueryRow("SELECT id FROM users WHERE email= $1", userEmail).Scan(
		&userMatch.Id)
	if err != nil {
		fmt.Println("email:", err)
		return err
	}

	return nil
}

func (db *UserDB) AddTokenDB(userid int, cookieToken string) error {
	query := `UPDATE users
	SET token = ?, expires = DATETIME('now', '+7 hours')
	WHERE ? = id` // expiration datetime = now + 1 hours
	if _, err := db.DB.Exec(query, cookieToken, userid); err != nil {
		return err
	}
	return nil
}

func (db *UserDB) RemoveTokenDB(token string) error {
	query := `UPDATE users
	SET token = NULL, expires = NULL
	WHERE token = ?`
	_, err := db.DB.Exec(query, token)
	return err
}

func (db *UserDB) GetUserByToken(token string) (*t.User, error) {
	user := &t.User{}
	err := db.DB.QueryRow("SELECT id, email FROM users WHERE token= $1", token).Scan(
		&user.Id,
		&user.Email)
	if err != nil {
		fmt.Println("GetUserByToken:   ", err)
		return nil, err
	}
	return user, nil
}
