package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/rs/zerolog/log"
	"sync"
	"todo-code-gen/internal/config"
	"todo-code-gen/internal/todo/entity"
)

var (
	once     sync.Once
	instance *sql.DB
	queries  entity.Queries
)

func ProvideDatabase(cfg config.DatabaseConfig) *entity.Queries {
	log.Printf("Try to create connection to database: %s:%s/%s with user: %s", cfg.Host, cfg.Port, cfg.DbName, cfg.User)

	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=preferred&timeout=5s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Err(err).Msg("Could not connect to mysql")
		}
		instance = db

		log.Printf("Connected to MySQL")
	})

	return entity.New(instance)
}
