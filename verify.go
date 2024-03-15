package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func verfiyCreateAd(c *gin.Context,ad *Ad) {
	if err := c.ShouldBindJSON(ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if (ad.Conditions.AgeStart < 0) != (ad.Conditions.AgeEnd < 0 ){// if both are -1, means no age condition
		c.JSON(http.StatusBadRequest, gin.H{"error": "AgeStart and AgeEnd should be positive"})
		return
	}
	if ad.Conditions.AgeStart > ad.Conditions.AgeEnd {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AgeStart should be less than AgeEnd"})
		return
	}
	startAt, err := time.Parse(time.RFC3339, ad.StartAt)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid StartAt format","data":ad.StartAt})
        return
    }

    endAt, err := time.Parse(time.RFC3339, ad.EndAt)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid EndAt format"})
        return
    }

    if startAt.After(endAt) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "StartAt should be less than EndAt"})
        return
    }
	if ad.Conditions.Gender!="" && ad.Conditions.Gender!="M" && ad.Conditions.Gender!="F"{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Gender should be M or F"})
		return
	}
	fmt.Println("Ad verified successfully")
}
