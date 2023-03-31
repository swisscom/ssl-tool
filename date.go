package main

import "time"

func formatDate(t time.Time) string {
	return t.In(time.Local).Format("2006-01-02 15:04:05")
}
