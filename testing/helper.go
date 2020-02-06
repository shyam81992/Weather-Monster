package testing

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Req struct {
	URI    string
	Method string
	Msg    []byte
}

// ErrMessage struct
type ErrMessage struct {
	Code       int    `json:"code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Error      string `json:"error"`
	Requesturl string `json:"requesturl"`
}

//RequestToWM method
func RequestToWM(req Req, args ...interface{}) *ErrMessage {
	var errmsg *ErrMessage
	var i = 0
	var retries = 2
	for i <= retries {

		var err *ErrMessage
		var response *http.Response
		var terr error
		switch strings.ToUpper(req.Method) {
		case "GET":
			response, terr = http.Get(req.URI)
		case "POST":
			response, terr = http.Post(req.URI, "application/json", bytes.NewBuffer(req.Msg))
		case "PATCH":
			client := &http.Client{}
			var hreq *http.Request
			hreq, terr = http.NewRequest("PATCH", req.URI, bytes.NewBuffer(req.Msg))
			if terr == nil {
				response, terr = client.Do(hreq)
			}
		case "DELETE":
			client := &http.Client{}
			var hreq *http.Request
			hreq, terr = http.NewRequest("DELETE", req.URI, nil)
			if terr == nil {
				response, terr = client.Do(hreq)
			}

		default:
			panic(errors.New("Invalid http method"))

		}

		if terr != nil {
			var errOBJ ErrMessage
			err = &errOBJ
			err.Error = "Error in processing the request"
			err.Message = terr.Error()
			err.Status = "400"
			err.Code = 400

		} else {
			// You can have your own error handling logic here
			data, terr := ioutil.ReadAll(response.Body)
			if terr != nil {
				var errOBJ ErrMessage
				err = &errOBJ
				err.Error = "Error reading the response"
				err.Message = "Error reading the response"
				err.Code = response.StatusCode
				err.Status = response.Status

			}

			if response.StatusCode != 200 {
				if json.Unmarshal(data, &err) != nil {
					var errOBJ ErrMessage
					err = &errOBJ
					err.Error = "Error response"
					err.Message = string(data)
					err.Code = response.StatusCode
					err.Status = response.Status
				} else {
					err.Code = response.StatusCode
					err.Status = response.Status
				}
			} else {
				if len(args) > 0 {
					if json.Unmarshal(data, &args[0]) != nil {
						var errOBJ ErrMessage
						err = &errOBJ
						err.Error = "Error response"
						err.Message = string(data)
						err.Code = response.StatusCode
						err.Status = response.Status
					}
				}

			}

		}
		if err != nil {
			if err.Code == 502 && i < retries {
				i++
				time.Sleep(2 * time.Second)
				continue
			}
		}
		errmsg = err
		i = retries + 1
	}

	return errmsg
}
