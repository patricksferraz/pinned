package utils

import (
	"os"
	"time"
)

func PString(str string) *string {
	return &str
}

func PFloat64(f float64) *float64 {
	return &f
}

func PInt(i int) *int {
	return &i
}

func PTime(t time.Time) *time.Time {
	return &t
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
