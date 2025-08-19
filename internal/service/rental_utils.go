package service

import (
	"CarRentalService/internal/models"
	"errors"
	"time"
)

var (
    ErrInvalidDateRange = errors.New("start date must be before end date")
	ErrParseDate = errors.New("invalid date")
	ErrPriceCalculate = errors.New("Price calculate error")
)


func checkDate(input models.StartRent) error {
	startTime, errStart := time.Parse("2006-01-02", input.StartDate)
	if errStart != nil {
		return ErrParseDate
	}
	endTime, errEnd := time.Parse("2006-01-02", input.EndDate)
	if errEnd != nil {
		return ErrParseDate
	}
	if !startTime.Before(endTime) {
        return ErrInvalidDateRange
    }
	return nil
}

func priceCalculation(start, end string, pricePerDay float64) (float64, error) {
	startTime, err := time.Parse("2006-01-02", start)
	endTime, err := time.Parse("2006-01-02", end)
	duration := endTime.Sub(startTime).Hours()/24
	return duration*pricePerDay, err
}