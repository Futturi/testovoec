package pkg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Hostname string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func InitDB(cfg Config) (*sqlx.DB, error) {
	con, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port =%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Hostname, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	return con, nil
}
