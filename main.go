package main

import (
	"log"
	"net/http"
	"strconv"
	_ "dcard_backend_assignment/docs" // replace with your project path
	"fmt"
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
// @Param   ad   body   main.Ad   true   "Create ad"
// @Success 201 {object} main.Ad
// @Router /api/v1/ad [post]
func adminApi(c *gin.Context) {
	var ad Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Created ad:")
	fmt.Fprintf(c.Writer, "Created ad: %v", ad)
	createAds(ad)
	c.JSON(http.StatusCreated, ad)
}

// @Summary Retieve Ad
// @Description Retieve ad by Query
// @Tags Ads
// @Accept   json
// @Produce  json
// @Param   offset     query   int      false   "Offset for pagination"
// @Param   limit      query   int      false   "Limit for pagination defalut 5"
// @Param   age        query   int      false   "Age to Query"
// @Param   gender     query   string   false   "Gender"
// @Param   country    query   string   false   "Country"
// @Param   platform   query   string   false   "Platform"
// @Success 200 {object} main.QueryCondition
// @Router /api/v1/ad [get]
func publicApi(c *gin.Context) {
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
	result, _ := retrieveAds(queryCondition)
	c.JSON(http.StatusOK, result)
}
func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.POST("/api/v1/ad", adminApi)
	router.GET("/api/v1/ad", publicApi)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := router.Run()
    if err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}