package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/models"
)

type WebHookCtl struct {
	Db db.IDB
}

func NewWebHookController(Db db.IDB) models.WebHookCtlInteface {
	return &WebHookCtl{Db: Db}
}

// CreateWebHookTable creates City table if not exits
func (w *WebHookCtl) CreateWebHookTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS webhook (
		id SERIAL primary key NOT NULL,
		city_id Integer NOT NULL,
		callback_url VARCHAR(2083) NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := w.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating webhook table")
		fmt.Println(err.Error())
		panic(err)
	}

}

//CreateWebHooks function
func (w *WebHookCtl) CreateWebHooks(c *gin.Context) {

	var wh models.WebHook
	if err := c.ShouldBindJSON(&wh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := w.Db.QueryRowContext(ctx, `INSERT INTO webhook(city_id, callback_url) 
	SELECT $1,CAST($2 AS VARCHAR) WHERE EXISTS (
        SELECT 1 FROM city WHERE id=$1
    ) RETURNING id`, wh.CityID, wh.CallbackURL).Scan(&wh.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Record Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {

		c.JSON(http.StatusOK, wh)
	}

}

//DeleteWebHooks function
func (w *WebHookCtl) DeleteWebHooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var wh models.WebHook

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `Delete from webhook WHERE id = $1 RETURNING id,city_id,callback_url`
	err = w.Db.QueryRowContext(ctx, sqlStatement, id).Scan(&wh.ID, &wh.CityID, &wh.CallbackURL)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Record Not Found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, wh)

}
