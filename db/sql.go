package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/shyam81992/Weather-Monster/config"
)

//go:generate mockgen -destination=./mock/db.go -package=mock github.com/shyam81992/Weather-Monster/db IDB,IRow

//go:generate mockgen -destination=./mock/sql.go -package=mock database/sql Result

type IDB interface {
	QueryRowContext(context.Context, string, ...interface{}) IRow
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

// IRow defines the methods that belong to a Row object.
type IRow interface {
	Scan(dest ...interface{}) error
}

type DB struct {
	Db *sql.DB
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) IRow {
	return db.Db.QueryRowContext(ctx, query, args...)
}
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.Db.ExecContext(ctx, query, args...)
}

// InitDb function
func InitDb() (IDB, error) {

	port, _ := strconv.Atoi(config.DbConfig["port"])
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbConfig["host"], port,
		config.DbConfig["user"], config.DbConfig["password"],
		config.DbConfig["dbname"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	fmt.Println("Successfully connected to db")
	return &DB{Db: db}, err

}
