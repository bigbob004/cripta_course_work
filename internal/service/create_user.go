package service

import (
	"cripta_course_work/internal/model"
	"cripta_course_work/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Возможно, нужен динамический salt
const (
	salt      = "sglsjkgesnglnskgslgn"
	signedKey = "gdrhaharjta35252rdfwl"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User, questions []model.Question) error {
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	for i, _ := range questions {
		questions[i].UserID = id
	}
	_, err = s.repo.CreateQuestions(questions)
	if err != nil {
		return err
	}
	return nil
}

/*
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signedKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
*/
