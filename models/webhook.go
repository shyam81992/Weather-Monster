package models

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=./mock/webhook.go -package=mock github.com/shyam81992/Weather-Monster/models WebHookCtlInteface

// WebHook modal
type WebHook struct {
	ID          int64  `form:"id" json:"id"`
	CityID      int64  `form:"city_id" json:"city_id" binding:"required"`
	CallbackURL string `form:"callback_url" json:"callback_url" binding:"required"`
}

type WebHookCtlInteface interface {
	CreateWebHookTable()
	CreateWebHooks(*gin.Context)
	DeleteWebHooks(*gin.Context)
}
