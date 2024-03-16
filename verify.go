package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func verfiyCreateAd(c *gin.Context,ad *Ad) error {
	if err := c.ShouldBindJSON(ad); err != nil {
		return err
	}
	if ad.Title == "" {
		return fmt.Errorf("Title should not be empty")
	}
	if ad.StartAt == "" {
		return fmt.Errorf("StartAt should not be empty")
	}
	if ad.EndAt == "" {
		return fmt.Errorf("EndAt should not be empty")
	}
	startAt, err := time.Parse(time.RFC3339, ad.StartAt)
    if err != nil {
        return fmt.Errorf("Error parsing StartAt! Should be in yyyy-mm-ddTHH:MM:SSZ format!")
    }

    endAt, err := time.Parse(time.RFC3339, ad.EndAt)
    if err != nil {
        return fmt.Errorf("Error parsing EndAt! Should be in yyyy-mm-ddTHH:MM:SSZ format!")
    }

	if startAt.After(endAt) {
        return fmt.Errorf("StartAt should be less than EndAt")
    }

	if (ad.Conditions.AgeStart < 0) != (ad.Conditions.AgeEnd < 0 ){// if both are -1, means no age condition
		return fmt.Errorf("AgeStart and AgeEnd should be positive")
	}

	if ad.Conditions.AgeStart > ad.Conditions.AgeEnd {
		return fmt.Errorf("AgeStart should be less than AgeEnd")
	}

	if ad.Conditions.Gender!="" && ad.Conditions.Gender!="M" && ad.Conditions.Gender!="F"{
		return fmt.Errorf("Gender should be M or F")
	}
	return nil
}
