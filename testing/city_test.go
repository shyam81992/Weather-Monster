package testing

import (
	"encoding/json"
	"fmt"

	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/config"
	"github.com/shyam81992/Weather-Monster/models"
)

func TestCity(t *testing.T) {
	config.LoadConfig()

	randomstr := RString(5)
	t.Run("City creation", testCityCreation(gin.H{"name": "Berlin-" + randomstr, "latitude": 52.520008, "longitude": 13.404954}, 200))
	t.Run("Update City", testCitUpdation(gin.H{"name": "Madras-" + randomstr, "latitude": 13.0827, "longitude": 80.2707},
		gin.H{"name": "Chennai-" + randomstr, "latitude": 13.0827, "longitude": 80.2707},
		200))

	t.Run("City Deletion", testCityDeletion(gin.H{"name": "bangalore-" + randomstr, "latitude": 52.520008, "longitude": 13.404954}, 200))

	t.Run("City creation throw error on duplicate", testCityCreation(gin.H{"name": "Berlin-" + randomstr, "latitude": 52.520008, "longitude": 13.404954}, 403))
	

}

func testCityCreation(city gin.H, status int) func(*testing.T) {
	return func(t *testing.T) {
		msg, _ := json.Marshal(city)
		req := Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/cities",
			Method: "post",
			Msg:    msg,
		}

		reqerr := RequestToWM(req)

		if reqerr != nil && reqerr.Code != status {
			t.Error(fmt.Sprintf("Error in creating the city %v ", city))
		}
	}
}

func testCitUpdation(city gin.H, updatevalue gin.H, status int) func(*testing.T) {

	return func(t *testing.T) {
		msg, _ := json.Marshal(city)
		req := Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/cities",
			Method: "post",
			Msg:    msg,
		}
		var tcity models.City
		reqerr := RequestToWM(req, &tcity)

		if reqerr != nil {
			t.Error(fmt.Sprintf("Error in creating the city %v ", city))
		}

		msg, _ = json.Marshal(updatevalue)

		req = Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/cities/" + strconv.FormatInt(tcity.ID, 10),
			Method: "patch",
			Msg:    msg,
		}

		reqerr = RequestToWM(req)

		if reqerr != nil {
			t.Error(fmt.Sprintf("Error in updating the city %v ", city))
		}
	}

}

func testCityDeletion(city gin.H, status int) func(*testing.T) {

	return func(t *testing.T) {
		msg, _ := json.Marshal(city)
		req := Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/cities",
			Method: "post",
			Msg:    msg,
		}
		var tcity models.City
		reqerr := RequestToWM(req, &tcity)

		if reqerr != nil {
			t.Error(fmt.Sprintf("Error in creating the city %v ", city))
		}

		req = Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/cities/" + strconv.FormatInt(tcity.ID, 10),
			Method: "delete",
		}

		reqerr = RequestToWM(req)

		if reqerr != nil {
			t.Error(fmt.Sprintf("Error in deleting the city %v ", city))
		}
	}

}
