package controllers

import "github.com/shyam81992/Weather-Monster/models"

type Config struct {
	CityController    models.CityCtlInterface
	TemperatureCtl    models.TemperatureCtlInteface
	WebHookController models.WebHookCtlInteface
}

func Init(c *Config) {
	//city
	c.CityController.CreateCityTable()
	c.CityController.CreateCityDeleteTrigger()

	// city delete
	c.CityController.CreateCityDeleteTable()

	// Temperature
	c.TemperatureCtl.CreateTemperatureTable()

	//WebHooks
	c.WebHookController.CreateWebHookTable()
}
