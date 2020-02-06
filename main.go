package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/config"
	"github.com/shyam81992/Weather-Monster/controllers"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/models"
)

func main() {

	config.LoadConfig()
	db.InitDb()
	models.Init()

	r := gin.Default()

	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/cities", controllers.CreateCity)
	r.PATCH("/cities/:id", controllers.UpdateCity)
	r.DELETE("/cities/:id", controllers.DeleteCity)

	r.POST("/temperatures", controllers.CreateTemperature)

	r.GET("/forecasts/:city_id", controllers.GetForecasts)

	r.POST("/webhooks", controllers.CreateWebHooks)
	r.DELETE("/webhooks/:id", controllers.DeleteWebHooks)

	r.POST("/api", controllers.API1)
	r.POST("/api2", controllers.API1)

	r.Run(":" + config.AppConfig["port"])
}
