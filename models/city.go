package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Weather-Monster/db"
)

// City modal
type City struct {
	ID        int64   `form:"id" json:"id"`
	Name      string  `form:"name" json:"name" binding:"required"`
	Latitude  float64 `form:"latitude" json:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" json:"longitude" binding:"required"`
}

// CreateCityTable creates City table if not exits
func CreateCityTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS city (
		id SERIAL primary key NOT NULL,
		name varchar(45) NOT NULL,
		latitude numeric NOT NULL,
		longitude numeric NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW(),
		UNIQUE(name)
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city table")
		fmt.Println(err.Error())
		panic(err)
	}

}

// CreateCityDeleteTrigger creates delete trigger for city table
func CreateCityDeleteTrigger() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE OR REPLACE FUNCTION process_city_delete() RETURNS TRIGGER AS 
	$city_delete$
		BEGIN
			INSERT INTO city_deleted(cid, name, latitude, longitude, created_at) SELECT OLD.id,OLD.name, OLD.latitude, OLD.longitude, OLD.created_at;
			RETURN NULL;  
		END;
	$city_delete$ 
	LANGUAGE plpgsql;
	CREATE TRIGGER process_city_delete AFTER DELETE ON city
		FOR EACH ROW EXECUTE PROCEDURE process_city_delete();`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city delete trigger")
		fmt.Println(err.Error())
		//panic(err)
	}
}
