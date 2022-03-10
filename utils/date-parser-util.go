package utils

import (
	"gin-mongo-api/logger"
	"time"
)

const (
	IST_OFFSET time.Duration = time.Second * -19800
	SoD        string        = "2006-01-02"
	EoD        string        = "2006-01-02 15:04:05"
)

func ParseISTDate(dateString string, eod bool) (time.Time, error) {
	var date time.Time
	var dateFormat string
	var err error

	if dateString == "" {
		return date, nil
	}
	if eod {
		dateFormat = EoD
		dateString = dateString + " 23:59:59"
	} else {
		dateFormat = SoD
	}

	if date, err = time.Parse(dateFormat, dateString); err != nil {
		return date, err
	}
	ISTdate := convertToIST(date)

	logger.InfoLogger.Println("Date parsed and converted to IST")
	return ISTdate, err
}

func convertToIST(dateTime time.Time) time.Time {
	return dateTime.Add(IST_OFFSET)
}
