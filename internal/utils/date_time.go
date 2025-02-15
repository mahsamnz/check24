package utils

import "time"

func GetCurrentISODate() string {
	return time.Now().Format("2006-01-02T15:04:05")
}

func GetCurrentUTCISODate() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05")
}
