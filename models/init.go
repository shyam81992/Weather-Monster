package models

func Init() {
	//city
	CreateCityTable()
	CreateCityDeleteTrigger()

	// city delete
	CreateCityDeleteTable()

	// Temperature
	CreateTemperatureTable()

	//WebHooks
	CreateWebHookTable()
}
