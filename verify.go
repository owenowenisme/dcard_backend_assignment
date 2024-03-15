package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func verfiyCreateAd(c *gin.Context,ad *Ad) error {
	if err := c.ShouldBindJSON(ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err

	}
	if (ad.Conditions.AgeStart < 0) != (ad.Conditions.AgeEnd < 0 ){// if both are -1, means no age condition
		return fmt.Errorf("AgeStart and AgeEnd should be positive")
	}
	if ad.Conditions.AgeStart > ad.Conditions.AgeEnd {
		return fmt.Errorf("AgeStart should be less than AgeEnd")
	}
	startAt, err := time.Parse(time.RFC3339, ad.StartAt)
    if err != nil {
        return err
    }

    endAt, err := time.Parse(time.RFC3339, ad.EndAt)
    if err != nil {
        return err
    }

    if startAt.After(endAt) {
        return fmt.Errorf("StartAt should be less than EndAt")
    }
	if ad.Conditions.Gender!="" && ad.Conditions.Gender!="M" && ad.Conditions.Gender!="F"{
		return fmt.Errorf("Gender should be M or F")
	}
	fmt.Println("Ad verified successfully")
	return nil
}
