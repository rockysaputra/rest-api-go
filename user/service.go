package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

type jwtService struct{}

func (j *jwtService) generateToken(claims jwt.MapClaims) (string, error) {
	// Membuat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangani token dengan kunci rahasia untuk menghasilkan tanda tangan
	secretKey := []byte("secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	tokenService := &jwtService{}
	user.Name = input.Name
	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user.PasswordHash = string(passwordHashed)

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Role = "User"

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	getToken, err := tokenService.generateToken(claims)

	if err != nil {
		return user, err
	}

	user.Token = getToken

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	// mencocokan password input dengan password db

	emailInput := input.Email
	passwordInput := input.Password

	findUser, err := s.repository.FindEmail(emailInput)

	if err != nil {
		return findUser, err
	}
	if findUser.ID == 0 {
		return findUser, errors.New("User Not Found")
	}

	passwordFromDB := findUser.PasswordHash

	passwordCompare := bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(passwordInput))

	if passwordCompare != nil {
		fmt.Println("password salah")
		return findUser, passwordCompare
	}

	return findUser, nil

}
