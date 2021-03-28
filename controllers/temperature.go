package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/camqp"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/models"
)

type TemperatureCtl struct {
	Db db.IDB
	camqp.CAMQPinterface
}

func NewTemperatureController(Db db.IDB, ampq camqp.CAMQPinterface) models.TemperatureCtlInteface {
	return &TemperatureCtl{Db: Db, CAMQPinterface: ampq}
}

// CreateCityTable creates City table if not exits
func (t *TemperatureCtl) CreateTemperatureTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS temperature (
		id SERIAL primary key NOT NULL,
		city_id Integer NOT NULL,
		min numeric NOT NULL,
		max numeric NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := t.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city table")
		fmt.Println(err.Error())
		panic(err)
	}

}

//CreateTemperature function
func (t *TemperatureCtl) CreateTemperature(c *gin.Context) {

	var temp models.Temperature
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var timestamp time.Time
	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := t.Db.QueryRowContext(ctx, `INSERT INTO temperature(city_id, min, max) 
	SELECT $1, $2, $3 WHERE EXISTS (
        SELECT 1 FROM city WHERE id=$1
    ) RETURNING id,created_at`, temp.CityID, temp.Min, temp.Max).Scan(&temp.ID, &timestamp)
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
		temp.Timestamp = timestamp.Unix()

		msg, _ := json.Marshal(gin.H{
			"city_id":   temp.CityID,
			"min":       temp.Min,
			"max":       temp.Max,
			"timestamp": temp.Timestamp,
		})
		t.Publishmsg(msg)
		c.JSON(http.StatusOK, temp)
	}

}

//GetForecasts function
func (t *TemperatureCtl) GetForecasts(c *gin.Context) {

	city_id, err := strconv.Atoi(c.Param("city_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var forecast models.Forecast

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err = t.Db.QueryRowContext(ctx, `SELECT city_id, avg(min), avg(max), COUNT(city_id) FROM temperature
	WHERE city_id=$1 AND created_at >= NOW() - INTERVAL '24 HOURS'
	GROUP BY city_id`, city_id).Scan(&forecast.CityID, &forecast.Min, &forecast.Max, &forecast.Sample)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No Records Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, forecast)
	}

}
