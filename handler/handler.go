package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/controllers"
	"github.com/shyam81992/Weather-Monster/models"
)

type Handler struct {
	cityController    models.CityCtlInterface
	temperatureCtl    models.TemperatureCtlInteface
	webHookController models.WebHookCtlInteface
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R                 *gin.Engine
	CityController    models.CityCtlInterface
	TemperatureCtl    models.TemperatureCtlInteface
	WebHookController models.WebHookCtlInteface
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		cityController:    c.CityController,
		temperatureCtl:    c.TemperatureCtl,
		webHookController: c.WebHookController,
	} // currently has no properties
	r := c.R.Group("/")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/cities", h.cityController.CreateCity)
	r.PATCH("/cities/:id", h.cityController.UpdateCity)
	r.DELETE("/cities/:id", h.cityController.DeleteCity)

	r.POST("/temperatures", h.temperatureCtl.CreateTemperature)

	r.GET("/forecasts/:city_id", h.temperatureCtl.GetForecasts)

	r.POST("/webhooks", h.webHookController.CreateWebHooks)
	r.DELETE("/webhooks/:id", h.webHookController.DeleteWebHooks)

	r.POST("/api", controllers.API1)
	r.POST("/api2", controllers.API1)
}
