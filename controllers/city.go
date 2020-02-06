package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/models"
)

//CreateCity function
func CreateCity(c *gin.Context) {

	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT INTO city(name, latitude, longitude) 
	SELECT CAST($1 AS VARCHAR), $2, $3 WHERE NOT EXISTS (
        SELECT 1 FROM city WHERE name=$1
    ) RETURNING id`, city.Name, city.Latitude, city.Longitude).Scan(&city.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(403, gin.H{
				"error": "Record already exits",
			})
			return
		}
		if pqError, ok := err.(*pq.Error); ok {
			if pqError.Code == "23505" {
				c.JSON(403, gin.H{
					"error": "Record already exits",
				})
			} else {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
		} else {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

	} else {
		c.JSON(200, gin.H{
			"id":   city.ID,
			"name": city.Name,
			"lat":  city.Latitude,
			"long": city.Longitude,
		})
	}

}

//UpdateCity function
func UpdateCity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	city.ID = id

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `UPDATE city SET name = $2, latitude = $3, longitude = $4 WHERE id = $1;`
	res, err := db.Db.ExecContext(ctx, sqlStatement, id, city.Name, city.Latitude, city.Longitude)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(200, city)
	} else {
		c.JSON(404, gin.H{
			"error": "Record Not Found",
		})
	}

}

//DeleteCity function
func DeleteCity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var city = models.City{}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `Delete from city  WHERE id = $1 RETURNING id,name,latitude,longitude`
	err := db.Db.QueryRowContext(ctx, sqlStatement, id).Scan(&city.ID, &city.Name, &city.Latitude, &city.Longitude)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Record Not Found",
			})
		} else {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	if city.ID > 0 {
		c.JSON(200, city)
	} else {
		c.JSON(404, gin.H{
			"error": "Record Not Found",
		})
	}

}
