package main

import (
	"fmt"
	"time"
)

func formatDate(t time.Time) string {
	return t.In(time.Local).Format("2006-01-02 15:04:05")
}

func formatDateWithExpiry(t time.Time) string {
	date := formatDate(t)
	now := time.Now()
	if now.After(t) {
		days := int(now.Sub(t).Hours() / 24)
		return fmt.Sprintf("%s (expired %d days ago)", date, days)
	}
	days := int(t.Sub(now).Hours() / 24)
	return fmt.Sprintf("%s (%d days left)", date, days)
}
