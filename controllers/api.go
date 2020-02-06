package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Weather-Monster/models"
)

func API1(c *gin.Context) {
	var temp models.Temperature
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("fullath = %s  data = %+v\n", c.Request.URL, temp)

	c.JSON(200, temp)

}
