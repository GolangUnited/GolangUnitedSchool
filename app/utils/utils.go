package utils

import "time"

func GetIntPtr(n int64) *int64 {
	return &n
}

func GetStrPtr(s string) *string {
	return &s
}

func GetTimePtr(t time.Time) *time.Time {
	return &t
}
