package models

//go:generate mockgen -destination=./mock/city.go -package=mock github.com/shyam81992/Weather-Monster/models CityCtlInterface

import (
	"github.com/gin-gonic/gin"
)

// City modal
type City struct {
	ID        int64   `form:"id" json:"id"`
	Name      string  `form:"name" json:"name" binding:"required"`
	Latitude  float64 `form:"latitude" json:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" json:"longitude" binding:"required"`
}
type CityCtlInterface interface {
	CreateCityTable()
	CreateCityDeleteTrigger()
	CreateCityDeleteTable()
	CreateCity(*gin.Context)
	UpdateCity(*gin.Context)
	DeleteCity(*gin.Context)
}
