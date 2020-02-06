package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/config"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/models"
	"github.com/streadway/amqp"
)

//CreateWebHooks function
func CreateWebHooks(c *gin.Context) {

	var wh models.WebHook
	if err := c.ShouldBindJSON(&wh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT INTO webhook(city_id, callback_url) 
	SELECT $1,CAST($2 AS VARCHAR) WHERE EXISTS (
        SELECT 1 FROM city WHERE id=$1
    ) RETURNING id`, wh.CityID, wh.CallbackURL).Scan(&wh.ID)
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

		c.JSON(200, wh)
	}

}

func publishmsg(msg []byte) {
	conn, err := amqp.Dial(config.RabbitConfig["uri"])
	if err != nil {
		fmt.Println(err, "Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err, "Failed to open a channel")
	}
	defer ch.Close()
	err = ch.Publish(
		"",                               // exchange
		config.RabbitConfig["queuename"], // routing key
		false,                            // mandatory
		false,                            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})

	if err != nil {
		fmt.Println(err, "Failed to publish a message")
	}
}

//DeleteWebHooks function
func DeleteWebHooks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var wh models.WebHook

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `Delete from webhook WHERE id = $1 RETURNING id,city_id,callback_url`
	err := db.Db.QueryRowContext(ctx, sqlStatement, id).Scan(&wh.ID, &wh.CityID, &wh.CallbackURL)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Record Not Found",
			})
		} else {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	if wh.ID > 0 {
		c.JSON(200, wh)
	} else {
		c.JSON(404, gin.H{
			"error": "Record Not Found",
		})
	}

}
