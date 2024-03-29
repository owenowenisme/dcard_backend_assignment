package main

import (
	_ "dcard_backend_assignment/docs"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"os"
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
	err := verfiyCreateAd(c, &ad)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_,err=CreateAds(ad)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Ad created as Id:%v successfully",ad.Title)})
}

// @Summary Retieve Ad
// @Description Retieve ad by Query (Now time should be retrieve from user frontend, in this case, I use now time in server)
// @Tags Ads
// @Produce  json
// @Param   offset     query   int      false   "Offset for pagination"
// @Param   limit      query   int      false   "Limit for pagination default 5"
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

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	HOST_URL := os.Getenv("HOST_URL")
	gin.SetMode(gin.DebugMode)
	//router := gin.Default()
	router := gin.New()
	router.POST("/api/v1/ad", AdminApi)
	router.GET("/api/v1/ad", PublicApi)
	url := ginSwagger.URL(HOST_URL+"/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := router.Run()
    if err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
