package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(connStr string) *Postgres {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to open postgres connection: %v", err))
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping postgres: %v", err))
	}

	fmt.Println("postgres connected successfully")
	return &Postgres{DB: db}
}

func (p *Postgres) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}
