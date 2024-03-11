package main

import (
	"net/http"
	"strconv"

	_ "dcard_backend_assignment/docs" // replace with your project path

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Sum two numbers
// @Description It adds two numbers
// @Tags numbers
// @Accept  json
// @Produce  json
// @Param   number1     path    int     true        "Number 1 to add"
// @Param   number2     path    int     true        "Number 2 to add"
// @Success 200 {object} Result
// @Router /sum/{number1}/{number2} [get]
func SumTwoNumbers(c *gin.Context) {
    number1, _ := strconv.Atoi(c.Param("number1"))
    number2, _ := strconv.Atoi(c.Param("number2"))
    sum := number1 + number2

    c.JSON(http.StatusOK, Result{Sum: sum})
}
// @Summary Create OMGGG
// @Description create ad by JSON
// @Tags ads
// @Accept  json
// @Produce  json
// @Param ad body main.Ad true "Create ad"
// @Success 200 {object} main.Ad
// @Router /api/v1/ad [post]
func adminApi(c *gin.Context){
    var ad Ad
    if err := c.ShouldBindJSON(&ad); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return 
    }
    c.JSON(http.StatusOK, ad)
}
// func publicApi(c *gin.Context){
    
// }
func main() {
    router := gin.Default()

    router.GET("/sum/:number1/:number2", SumTwoNumbers)
    router.POST("/api/v1/ad", adminApi)
    // router.GET("/api/v1/ad", publicApi)
    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    router.Run()
}