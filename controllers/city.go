package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/shyam81992/Weather-Monster/db"
	"github.com/shyam81992/Weather-Monster/models"
)

type CityCtl struct {
	Db db.IDB
}

func NewCityController(Db db.IDB) models.CityCtlInterface {
	return &CityCtl{Db: Db}
}

// CreateCityTable creates City table if not exits
func (cc *CityCtl) CreateCityTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS city (
		id SERIAL primary key NOT NULL,
		name varchar(45) NOT NULL,
		latitude numeric NOT NULL,
		longitude numeric NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW(),
		UNIQUE(name)
	  )`
	_, err := cc.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city table")
		fmt.Println(err.Error())
		panic(err)
	}

}

// CreateCityDeleteTrigger creates delete trigger for city table
func (cc *CityCtl) CreateCityDeleteTrigger() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE OR REPLACE FUNCTION process_city_delete() RETURNS TRIGGER AS 
	$city_delete$
		BEGIN
			INSERT INTO city_deleted(cid, name, latitude, longitude, created_at) SELECT OLD.id,OLD.name, OLD.latitude, OLD.longitude, OLD.created_at;
			RETURN NULL;  
		END;
	$city_delete$ 
	LANGUAGE plpgsql;
	CREATE TRIGGER process_city_delete AFTER DELETE ON city
		FOR EACH ROW EXECUTE PROCEDURE process_city_delete();`
	_, err := cc.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city delete trigger")
		fmt.Println(err.Error())
		//panic(err)
	}
}

//CreateCityDeleteTable Create City table if not exits
func (cc *CityCtl) CreateCityDeleteTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS city_deleted (
		id SERIAL primary key NOT NULL,
		cid Integer NOT NULL,
		name varchar(45) NOT NULL,
		latitude numeric NOT NULL,
		longitude numeric NOT NULL,
		created_at timestamptz NOT NULL,
		deleted_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := cc.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating city_delete table")
		fmt.Println(err.Error())
		panic(err)
	}

}

//CreateCity function
func (cc *CityCtl) CreateCity(c *gin.Context) {

	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := cc.Db.QueryRowContext(ctx, `INSERT INTO city(name, latitude, longitude) 
	SELECT CAST($1 AS VARCHAR), $2, $3 WHERE NOT EXISTS (
        SELECT 1 FROM city WHERE name=$1
    ) RETURNING id`, city.Name, city.Latitude, city.Longitude).Scan(&city.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Record already exits",
			})
			return
		}
		if pqError, ok := err.(*pq.Error); ok {
			if pqError.Code == "23505" {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Record already exits",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
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
func (cc *CityCtl) UpdateCity(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	city.ID = id

	ctx := c.Request.Context()
	sqlStatement := `UPDATE city SET name = $2, latitude = $3, longitude = $4 WHERE id = $1;`
	res, err := cc.Db.ExecContext(ctx, sqlStatement, id, city.Name, city.Latitude, city.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, city)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record Not Found",
		})
	}

}

//DeleteCity function
func (cc *CityCtl) DeleteCity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var city = models.City{}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `Delete from city  WHERE id = $1 RETURNING id,name,latitude,longitude`
	err = cc.Db.QueryRowContext(ctx, sqlStatement, id).Scan(&city.ID, &city.Name, &city.Latitude, &city.Longitude)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Record Not Found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, city)

}
