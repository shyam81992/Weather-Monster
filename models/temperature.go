package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Weather-Monster/db"
)

// Temperature modal
type Temperature struct {
	ID        int64   `form:"id" json:"id"`
	CityID    int64   `form:"city_id" json:"city_id" binding:"required"`
	Max       float64 `form:"max" json:"max" binding:"required"`
	Min       float64 `form:"min" json:"min" binding:"required"`
	Timestamp int64   `form:"timestamp" json:"timestamp"`
}

// Forecast modal
type Forecast struct {
	CityID int64   `form:"city_id" json:"city_id"`
	Max    float64 `form:"max" json:"max"`
	Min    float64 `form:"min" json:"min"`
	Sample int64   `form:"sample" json:"sample"`
}

// CreateCityTable creates City table if not exits
func CreateTemperatureTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS temperature (
		id SERIAL primary key NOT NULL,
		city_id Integer NOT NULL,
		min numeric NOT NULL,
		max numeric NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city table")
		fmt.Println(err.Error())
		panic(err)
	}

}
