package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/noyandey/go-ddd-starter/config"
)

func GetConnectionString(cfg *config.DBConfig) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	if !cfg.EnableSSLMode {
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnection(cfg *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cfg)
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
