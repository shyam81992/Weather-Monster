package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/models"
)

//CreateTemperature function
func CreateTemperature(c *gin.Context) {

	var temp models.Temperature
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var timestamp time.Time
	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT INTO temperature(city_id, min, max) 
	SELECT $1, $2, $3 WHERE EXISTS (
        SELECT 1 FROM city WHERE id=$1
    ) RETURNING id,created_at`, temp.CityID, temp.Min, temp.Max).Scan(&temp.ID, &timestamp)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Record Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		temp.Timestamp = timestamp.Unix()

		msg, _ := json.Marshal(gin.H{
			"city_id" : temp.CityID,
			"min" : temp.Min,
			"max" : temp.Max,
			"timestamp" : temp.Timestamp,
		})
		publishmsg(msg)
		c.JSON(200, temp)
	}

}

//GetForecasts function
func GetForecasts(c *gin.Context) {

	city_id, _ := strconv.Atoi(c.Param("city_id"))
	var forecast models.Forecast

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `SELECT city_id, avg(min), avg(max), COUNT(city_id) FROM temperature
	WHERE city_id=$1 GROUP BY city_id`, city_id).Scan(&forecast.CityID, &forecast.Min, &forecast.Max, &forecast.Sample)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Record Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, forecast)
	}

}
