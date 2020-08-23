package utils

import (
	"time"
)

// MakeTimestampMilli return time now in milli second
func MakeTimestampMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
