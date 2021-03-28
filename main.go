package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/camqp"
	"github.com/shyam81992/Weather-Monster/config"
	"github.com/shyam81992/Weather-Monster/controllers"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/handler"
)

func main() {

	config.LoadConfig()
	Db, _ := db.InitDb()

	campq := camqp.CAMQP{}

	cityController := controllers.NewCityController(Db)
	temperatureController := controllers.NewTemperatureController(Db, &campq)
	webHookController := controllers.NewWebHookController(Db)

	controllers.Init(&controllers.Config{
		CityController:    cityController,
		TemperatureCtl:    temperatureController,
		WebHookController: webHookController,
	})

	r := gin.Default()

	r.Use(gin.Recovery())

	handler.NewHandler(&handler.Config{
		R:                 r,
		CityController:    cityController,
		TemperatureCtl:    temperatureController,
		WebHookController: webHookController,
	})

	r.Run(":" + config.AppConfig["port"])
}
