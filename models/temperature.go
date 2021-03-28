package models

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=./mock/temperature.go -package=mock github.com/shyam81992/Weather-Monster/models TemperatureCtlInteface

// Temperature modal
type Temperature struct {
	ID        int64   `form:"id" json:"id"`
	CityID    int64   `form:"city_id" json:"city_id" binding:"required"`
	Max       float64 `form:"max" json:"max" binding:"required"`
	Min       float64 `form:"min" json:"min" binding:"required"`
	Timestamp int64   `form:"timestamp" json:"timestamp"`
}

// Forecast modal
type Forecast struct {
	CityID int64   `form:"city_id" json:"city_id"`
	Max    float64 `form:"max" json:"max"`
	Min    float64 `form:"min" json:"min"`
	Sample int64   `form:"sample" json:"sample"`
}

type TemperatureCtlInteface interface {
	CreateTemperatureTable()
	CreateTemperature(*gin.Context)
	GetForecasts(*gin.Context)
}
