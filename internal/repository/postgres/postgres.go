package postgres

import (
	"fmt"

	"github.com/Levap123/task-manager-auth-service/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDb(cfg *config.Configs) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=localhost port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		 cfg.DB_PORT, cfg.DB_USER_NAME, cfg.DB_PASSWORD, cfg.DB_NAME)
	db, err := sqlx.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := createTables(db); err != nil {
		return nil, err
	}
	return db, nil
}
