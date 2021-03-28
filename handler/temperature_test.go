package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	Camqp "github.com/shyam81992/Weather-Monster/camqp"
	camqp "github.com/shyam81992/Weather-Monster/camqp/mock"
	"github.com/shyam81992/Weather-Monster/controllers"
	DB "github.com/shyam81992/Weather-Monster/db"
	db "github.com/shyam81992/Weather-Monster/db/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTemperature(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDb := db.NewMockIDB(mockCtrl)
	mockRow := db.NewMockIRow(mockCtrl)
	mockCamqp := camqp.NewMockCAMQPinterface(mockCtrl)

	cityController := &controllers.CityCtl{Db: mockDb}

	temperatureCtl := &controllers.TemperatureCtl{Db: mockDb, CAMQPinterface: mockCamqp}

	webHookController := &controllers.WebHookCtl{Db: mockDb}

	if os.Getenv("integration_testing") == "true" {
		db, _ := DB.InitDb()
		campq := Camqp.CAMQP{}
		cityController = &controllers.CityCtl{Db: db}

		temperatureCtl = &controllers.TemperatureCtl{Db: db, CAMQPinterface: &campq}

		webHookController = &controllers.WebHookCtl{Db: db}
	}

	NewHandler(&Config{
		R:                 router,
		CityController:    cityController,
		TemperatureCtl:    temperatureCtl,
		WebHookController: webHookController,
	})

	var testCases = []struct {
		name          string
		input         gin.H
		buildStubs    func(mockDb *db.MockIDB, mockRow *db.MockIRow)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "Insert Record Success",
			input: gin.H{
				"city_id": 1,
				"max":     35,
				"min":     32,
			},
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
				mockDb.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow)
				mockRow.EXPECT().Scan(gomock.Any()).Return(nil).DoAndReturn(func(val ...interface{}) (err error) {
					var arr [][]byte
					//var date = time.Now()
					arr = append(arr, []byte("1"), []byte(`"2009-11-10T23:00:00Z"`))
					for i := 0; i < len(val); i++ {
						err = json.Unmarshal(arr[i], val[i])
						if err != nil {
							break
						}
					}
					return err
				})
				mockCamqp.EXPECT().Publishmsg(gomock.Any()).Return()
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Bad request data",
			input: gin.H{
				"city_id": 1,
				"max":     35,
			},
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Data Conflict level",
			input: gin.H{
				"city_id": 1,
				"max":     35,
				"min":     32,
			},
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
				mockDb.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow)
				mockRow.EXPECT().Scan(gomock.Any()).Return(errors.New("sql: no rows in result set"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
		{
			name: "Internal server ERROR level",
			input: gin.H{
				"city_id": 1,
				"max":     35,
				"min":     32,
			},
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
				mockDb.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow)
				mockRow.EXPECT().Scan(gomock.Any()).Return(errors.New("Internal server error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			//.After(callfirst)
			// create a request body with invalid fields
			reqBody, err := json.Marshal(test.input)
			assert.NoError(t, err)

			if os.Getenv("integration_testing") != "true" {
				test.buildStubs(mockDb, mockRow)
			}

			request, err := http.NewRequest(http.MethodPost, "/temperatures", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)

			test.checkResponse(t, rr)

		})
	}

}

func TestGetForecasts(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDb := db.NewMockIDB(mockCtrl)
	mockRow := db.NewMockIRow(mockCtrl)
	cityController := &controllers.CityCtl{Db: mockDb}

	temperatureCtl := &controllers.TemperatureCtl{Db: mockDb}

	webHookController := &controllers.WebHookCtl{Db: mockDb}

	if os.Getenv("integration_testing") == "true" {
		db, _ := DB.InitDb()
		campq := Camqp.CAMQP{}
		cityController = &controllers.CityCtl{Db: db}

		temperatureCtl = &controllers.TemperatureCtl{Db: db, CAMQPinterface: &campq}

		webHookController = &controllers.WebHookCtl{Db: db}
	}

	NewHandler(&Config{
		R:                 router,
		CityController:    cityController,
		TemperatureCtl:    temperatureCtl,
		WebHookController: webHookController,
	})

	var testCases = []struct {
		name          string
		input         string
		buildStubs    func(mockDb *db.MockIDB, mockRow *db.MockIRow)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:  "Get Record Success",
			input: "1",
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
				mockDb.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow).Times(1)
				mockRow.EXPECT().Scan(gomock.Any()).DoAndReturn(func(val ...interface{}) (err error) {
					var arr [][]byte
					arr = append(arr, []byte("1"), []byte("10"), []byte("10"), []byte("5"))
					for i := 0; i < len(val); i++ {
						err = json.Unmarshal(arr[i], val[i])
						if err != nil {
							break
						}
					}
					return err
				})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "Bad request data",
			input: "4sfs",
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:  "Resource Not Found",
			input: "1",
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
				mockDb.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow)
				mockRow.EXPECT().Scan(gomock.Any()).Return(errors.New("sql: no rows in result set"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},

		{
			name:  "Internal server ERROR",
			input: "1",
			buildStubs: func(mockDb *db.MockIDB, mockRow *db.MockIRow) {
				mockDb.EXPECT().QueryRowContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRow)
				mockRow.EXPECT().Scan(gomock.Any()).Return(errors.New("Internal Server Error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			//.After(callfirst)
			// create a request body with invalid fields
			reqBody, err := json.Marshal(test.input)
			assert.NoError(t, err)
			if os.Getenv("integration_testing") != "true" {
				test.buildStubs(mockDb, mockRow)
			}
			request, err := http.NewRequest(http.MethodGet, "/forecasts/"+test.input, bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)

			test.checkResponse(t, rr)

		})
	}

}
