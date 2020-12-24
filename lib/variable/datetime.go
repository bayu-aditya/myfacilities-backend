package variable

import (
	"time"
)

// DateTimeNow with output is pointer
func DateTimeNow() *time.Time {
	now := time.Now()
	return &now
}
