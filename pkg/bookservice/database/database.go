package database

import (
	"database/sql"

	"github.com/elboboua/bookmark/pkg/bookservice"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
)

type IDatabaseDriver interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type IDatabase interface {
	GetAllBooks() ([]bookservice.Book, error)
}

func CreateNewDatabase () IDatabase {
	sqlite_file := "db/book.db"
	db, err := sql.Open("sqlite3", sqlite_file)
	if err != nil {
		panic("unable to open db: "+ err.Error())
	}

	// run migrations
	if err := goose.SetDialect("sqlite3"); err != nil {
		panic("unable to set goose dialect: "+err.Error())
	}
	if err := goose.Up(db, "db/migrations"); err != nil {
		panic("unable to run goose migration: "+err.Error())
	}

	// set connection configs
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	sqliteConnection := Sqlite{
		db: db,
	}

	return &sqliteConnection
}