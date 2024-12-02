package connection

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"rijik.id/restapi_gofiber/internal/config"
)

var DB *pgx.Conn

func InitDB() {
	connStr := "postgres://"+config.DBUser+":"+config.DBPassword+"@"+config.DBHost+":"+config.DBPort+"/"+config.DBName
	var err error
	DB, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("gagal terhubung ke database: %v\n", err)
	}
	log.Println("Berhasil terhubung ke database")
}
