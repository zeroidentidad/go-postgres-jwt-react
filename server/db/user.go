package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Register struct
type Register struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ResetPassword struct {
	ID              int    `json:"id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// Login struct
type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateReset struct {
	Email string `json:"email"`
}

//User struct
type User struct {
	//ID string
	Password  string `json:"password"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var cost = 8

//HashPassword hace hash la contraseña de usuario
func HashPassword(user *Register) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
}

func CreateHashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

//CheckPasswordHash compara hash con password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
