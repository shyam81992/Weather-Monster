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

func TestWebHook(t *testing.T) {
	config.LoadConfig()

	t.Run("WebHook creation", testWebHookCreation(gin.H{
		"city_id":      1,
		"callback_url": "http://localhost:8080/api",
	}, 200))

	t.Run("WebHook creation", testWebHookDeletion(gin.H{
		"city_id":      2,
		"callback_url": "http://localhost:8080/api",
	}, 200))

}

func testWebHookCreation(temp gin.H, status int) func(*testing.T) {
	return func(t *testing.T) {
		msg, _ := json.Marshal(temp)
		req := Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/webhooks",
			Method: "post",
			Msg:    msg,
		}

		reqerr := RequestToWM(req)
		//fmt.Println(fmt.Sprintf("Error in creating the city %v  error %v", temp, reqerr))
		if reqerr != nil && reqerr.Code != status {
			t.Error(fmt.Sprintf("Error in creating the webhook %v  error %v", temp, reqerr))
		}
	}
}


func testWebHookDeletion(temp gin.H, status int) func(*testing.T) {
	return func(t *testing.T) {

		var wh  models.WebHook
		msg, _ := json.Marshal(temp)
		req := Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/webhooks",
			Method: "post",
			Msg:    msg,
		}

		reqerr := RequestToWM(req, &wh)
		//fmt.Println(fmt.Sprintf("Error in creating the city %v  error %v", temp, reqerr))
		if reqerr != nil && reqerr.Code != status {
			t.Error(fmt.Sprintf("Error in creating the webhook %v  error %v", temp, reqerr))
		}

		req = Req{
			URI:    "http://" + config.AppConfig["host"] + ":" + config.AppConfig["port"] + "/webhooks/" + strconv.FormatInt(wh.ID, 10),
			Method: "delete",
		}

		reqerr = RequestToWM(req)

		if reqerr != nil {
			t.Error(fmt.Sprintf("Error in deleting the webhook %v ", wh))
		}
	}
}