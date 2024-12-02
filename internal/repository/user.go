package repository

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx"
	"github.com/oklog/ulid"
	"rijik.id/restapi_gofiber/domain"
	"rijik.id/restapi_gofiber/internal/connection"
)

func CreateUser(user domain.User) error {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	user.ID = ulid.MustNew(ulid.Now(), entropy).String()

	_, err := connection.DB.Exec(context.Background(), "INSERT INTO users(id, username, email, password) VALUES($1, $2, $3, $4)",
		user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("error menambahkan user ke database.users: %v", err)
		return err
	}

	return nil
}

func GetUserByUsername(username string) (domain.User, error) {
	var user domain.User

	log.Printf("Querying user by username: %s", username)

	row := connection.DB.QueryRow(context.Background(), "SELECT id, username, email, password FROM users WHERE username = $1", username)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {

			return domain.User{}, nil
		}
		log.Printf("Error retrieving user by username: %v", err)
		return user, err
	}

	log.Printf("User found: %v", user)
	return user, nil
}

func GetUserByEmail(email string) (domain.User, error) {
	var user domain.User

	log.Printf("Querying user by email: %s", email)

	row := connection.DB.QueryRow(context.Background(), "SELECT id, username, email, password FROM users WHERE email = $1", email)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {

			return domain.User{}, nil
		}
		log.Printf("Error retrieving user by email: %v", err)
		return user, err
	}

	return user, nil
}

func GetUserByID(userID string) (domain.User, error) {
	var user domain.User

	row := connection.DB.QueryRow(context.Background(), "SELECT id, username, email, password FROM users WHERE id = $1", userID)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}
