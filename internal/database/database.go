package database

import (
	"fmt"
	"go-away-2024/internal/config"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db := sqlx.MustConnect(cfg.Database.Driver, dsn)
	log.Infof("Connected to database: %s", dsn)
	return db
}
