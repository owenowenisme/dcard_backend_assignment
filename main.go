package main

import (
	_ "dcard_backend_assignment/docs" // replace with your project path
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Create Ad
// @Description create ad by JSON
// @Tags Ads
// @Accept   json
// @Produce  json
// @Param   ad   body   main.Ad   true  "Create ad"
// @Success 201 {object} main.Ad
// @Router /api/v1/ad [post]
func AdminApi(c *gin.Context) {
	ad := InitAds()
	verfiyCreateAd(c, &ad)
	id,err:=CreateAds(ad)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Ad created as Id:%d successfully",id)})
}

// @Summary Retieve Ad
// @Description Retieve ad by Query (Now time should be retrieve from user frontend, in this case, I use now time in server)
// @Tags Ads
// @Produce  json
// @Param   offset     query   int      false   "Offset for pagination"
// @Param   limit      query   int      false   "Limit for pagination defalut 5"
// @Param   age        query   int      false   "Age to Query"
// @Param   gender     query   string   false   "Gender"
// @Param   country    query   string   false   "Country"
// @Param   platform   query   string   false   "Platform"
// @Success 200 {object} main.QueryCondition
// @Router /api/v1/ad [get]
func PublicApi(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
    limit, _ := strconv.Atoi(c.Query("limit"))
	age, _ := strconv.Atoi(c.Query("age"))
    gender := c.Query("gender")
	country := c.Query("country")
	platform := c.Query("platform")

	queryCondition := QueryCondition{
		Offset:   offset,
        Limit:    limit,
		Age:      age,
        Gender:   gender,
		Country:  country,
		Platform: platform,
	}
	result, err := RetrieveAds(queryCondition)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(result) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No data found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Get now time in DB
// @Description Now
// @Tags For development
// @Accept   json
// @Produce  json
// @Success 200 {object} string
// @Router /api/v1/now [get]
func NowTimeInDB(c *gin.Context) {
	result := time.Now().Local().Format("2006-01-02T15:04:05Z")
	c.JSON(http.StatusOK, result)
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.POST("/api/v1/ad", AdminApi)
	router.GET("/api/v1/ad", PublicApi)
	router.GET("/api/v1/now", NowTimeInDB)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := router.Run()
    if err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
