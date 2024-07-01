package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"online-shop/internal/config"
)

func ConnectionPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",

		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err = sqlx.Open("postgres", dataSource)

	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifeConnection) * time.Second)
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))

	return
}
