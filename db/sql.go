package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/shyam81992/Weather-Monster/config"
)

//Db connection
var Db *sql.DB

// InitDb function
func InitDb() (err error) {

	port, _ := strconv.Atoi(config.DbConfig["port"])
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbConfig["host"], port,
		config.DbConfig["user"], config.DbConfig["password"],
		config.DbConfig["dbname"])

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	Db.SetConnMaxLifetime(time.Hour)


	fmt.Println("Successfully connected to db")
	return nil

}
