package testing

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/config"
)

func TestTemperature(t *testing.T) {
	config.LoadConfig()

	t.Run("Temperature creation", testTemperatureCreation(gin.H{
		"city_id": 1,
		"max":     35,
		"min":     32,
	}, 200))
	t.Run("Temperature creation need to throw error on city not found", testTemperatureCreation(gin.H{
		"city_id": 198908,
		"max":     35,
		"min":     32,
	}, 404))

}

func TestForecasts(t *testing.T) {
	config.LoadConfig()

	t.Run("Forecast", testForcastRequest(1, 200))

	t.Run("Forecast doesn't return records", testForcastRequest(10000, 404))

}

func testTemperatureCreation(temp gin.H, status int) func(*testing.T) {
	return func(t *testing.T) {
		msg, _ := json.Marshal(temp)
		req := Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/temperatures",
			Method: "post",
			Msg:    msg,
		}

		reqerr := RequestToWM(req)
		//fmt.Println(fmt.Sprintf("Error in creating the city %v  error %v", temp, reqerr))
		if reqerr != nil && reqerr.Code != status {
			t.Error(fmt.Sprintf("Error in creating the temperature %v  error %v", temp, reqerr))
		}
	}
}

func testForcastRequest(city_id int, status int) func(*testing.T) {
	return func(t *testing.T) {

		req := Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/forecasts/" + strconv.Itoa(city_id),
			Method: "get",
		}

		reqerr := RequestToWM(req)
		//fmt.Println(fmt.Sprintf("Error in creating the city %v  error %v", temp, reqerr))
		if reqerr != nil && reqerr.Code != status {
			t.Error(fmt.Sprintf("Error in forcast api %v  error %v", req.URI, reqerr))
		}
	}
}
