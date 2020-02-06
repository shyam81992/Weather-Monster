package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Weather-Monster/db"
)

//CreateCityDeleteTable Create City table if not exits
func CreateCityDeleteTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS city_deleted (
		id SERIAL primary key NOT NULL,
		cid Integer NOT NULL,
		name varchar(45) NOT NULL,
		latitude numeric NOT NULL,
		longitude numeric NOT NULL,
		created_at timestamptz NOT NULL,
		deleted_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city_delete table")
		fmt.Println(err.Error())
		panic(err)
	}

}
