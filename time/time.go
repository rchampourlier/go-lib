package time

import (
	"time"
)

// MsToTime converts an epoch time in milliseconds to a
// `time.Time`
func MsToTime(ms int64) (time.Time, error) {
	return time.Unix(0, ms*int64(time.Millisecond)), nil
}
