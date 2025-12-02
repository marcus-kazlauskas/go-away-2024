package database

import (
	"fmt"
	"go-away-2024/internal/config"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DatabaseCfg.Host,
		config.DatabaseCfg.Port,
		config.DatabaseCfg.User,
		config.DatabaseCfg.Password,
		config.DatabaseCfg.Name,
		config.DatabaseCfg.SslMode,
	)

	DB = sqlx.MustConnect(config.DatabaseCfg.Driver, dsn)
	log.Infof("Connected to database: %s", dsn)
}
