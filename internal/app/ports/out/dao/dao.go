package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang-clean-architecture/internal/app/ports/out/dao/repository/note"
	"golang-clean-architecture/pkg"
	"os"
)

type Dao struct {
	NoteRepo note.Repository
}

func newDbConnection(config *pkg.AppConfig) *sqlx.DB {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.DataBaseName,
	)
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db
}

func Create(config *pkg.AppConfig) (*Dao, *sqlx.DB) {
	db := newDbConnection(config)
	return &Dao{
		NoteRepo: *note.New(db),
	}, db
}
