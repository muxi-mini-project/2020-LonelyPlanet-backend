package util

import (
	"strconv"
	"time"
)

func NextDay() string {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	result := strconv.FormatInt(next.Unix(), 10)
	return result
}

func LastWeek() string {
	now := time.Now()
	next := now.AddDate(0, 0, -7)
	result := strconv.FormatInt(next.Unix(), 10)
	return result
}

func LastMonth() string {
	now := time.Now()
	next := now.AddDate(0, -1, 0)
	result := strconv.FormatInt(next.Unix(), 10)
	return result
}
