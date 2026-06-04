package helper

import (
	"fmt"
	"time"
)


func FormatTimeIntoTimeElapsed(t *time.Time) *string {
	if t == nil {
		return nil
	}
	
	duration := time.Since(*t)
	if duration.Hours() < 24 {
		hours := int(duration.Hours())
		if hours == 1 {
			result := "1 hour ago"
			return &result
		}
		result := fmt.Sprintf("%d hours ago", hours)
		return &result
	}
	days := int(duration.Hours() / 24)
	if days == 1 {
		result := "1 day ago"
		return &result
	}
	result := fmt.Sprintf("%d days ago", days)
	return &result
}
