package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Weather-Monster/db"
)

// WebHook modal
type WebHook struct {
	ID          int64  `form:"id" json:"id"`
	CityID      int64  `form:"city_id" json:"city_id" binding:"required"`
	CallbackURL string `form:"callback_url" json:"callback_url" binding:"required"`
}

// CreateWebHookTable creates City table if not exits
func CreateWebHookTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS webhook (
		id SERIAL primary key NOT NULL,
		city_id Integer NOT NULL,
		callback_url VARCHAR(2083) NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating webhook table")
		fmt.Println(err.Error())
		panic(err)
	}

}
