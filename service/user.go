package service

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"majoo-backend-test/constant"
	"majoo-backend-test/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *Service) Login(userName, password string) (*model.UserLogin, int, error) {

	account, err := s.repository.GetAccount(userName)
	if err == sql.ErrNoRows {
		return nil, http.StatusNotFound, errors.New(constant.MSG_ERROR_USER_NOT_FOUND)
	}
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(constant.MSG_ERROR_DATABASE)
	}

	hash := md5.Sum([]byte(password))
	passwordMD5 := hex.EncodeToString(hash[:])

	if passwordMD5 != account.Password {
		return nil, http.StatusUnauthorized, errors.New(constant.MSG_ERROR_INVALID_USERNAME_PASSWORD)
	}

	claims := &constant.JwtCustomClaims{
		ID:       account.ID,
		UserName: account.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		log.Println(err.Error())
		return nil, http.StatusInternalServerError, errors.New(constant.MSG_ERROR_TOKEN)
	}

	userLogin := &model.UserLogin{
		Token: t,
		User: model.User{
			ID:       account.ID,
			UserName: account.UserName,
			Name:     account.Name,
		},
	}

	return userLogin, 0, nil
}
