package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"rijik.id/restapi_gofiber/domain"
	"rijik.id/restapi_gofiber/dto"
	"rijik.id/restapi_gofiber/internal/config"
	"rijik.id/restapi_gofiber/internal/repository"
	"rijik.id/restapi_gofiber/internal/utils"
)

func RegisterUser(userDTO dto.UserRegisterDTO) error {

	user, err := repository.GetUserByUsername(userDTO.Username)
	if err != nil {
		return err
	}
	if user.Username != "" {

		return errors.New("username already exists")
	}

	userByEmail, err := repository.GetUserByEmail(userDTO.Email)
	if err != nil {
		return err
	}
	if userByEmail.Email != "" {

		return errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(userDTO.Password)
	if err != nil {
		return err
	}

	user = domain.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: hashedPassword,
	}

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(userDTO dto.UserLoginDTO) (string, error) {

	user, err := repository.GetUserByUsername(userDTO.Username)
	if err != nil {
		return "", errors.New("pengguna tidak ditemukan")
	}

	if !utils.CheckPasswordHash(userDTO.Password, user.Password) {
		return "", errors.New("password anda salah")
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateJWT(userID string) (string, error) {

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tk, nil
}
