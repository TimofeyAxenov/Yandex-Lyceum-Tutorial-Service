package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	UserName string `env:"POSTGRES_USER" envDefault:"root"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"123"`
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	Dbname   string `env:"POSTGRES_DB" envDefault:"yandex"`
}

type DB struct {
	Db *sqlx.DB
}

func New(config Config) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.UserName, config.Password, config.Dbname, config.Host, config.Port)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Conn(context.Background()); err != nil {
		return nil, fmt.Errorf("Failed to connect to db: %w", err)
	}

	return &DB{Db: db}, nil
}
