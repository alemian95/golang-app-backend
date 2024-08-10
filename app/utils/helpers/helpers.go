package helpers

import "time"

func NewAssocArray() map[string]any {
	return make(map[string]any)
}

func GetCurrentDateTimestamp() string {
	return time.Now().Format("2006-01-02T15:04:05 +010000")
}
