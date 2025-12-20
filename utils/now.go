package utils

import "time"

func NowString() string {
	return time.Now().Format(time.RFC3339)
}
