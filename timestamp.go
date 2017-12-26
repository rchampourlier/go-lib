package golib

import (
	strftime "github.com/jehiah/go-strftime"
	"strings"
	"time"
)

// Timestamp returns a timestamp string for the specified time
// using a standart format: `%Y%m%dT%H%M%S%L%Z`.
func Timestamp(t time.Time) string {
	timestamp := strftime.Format("%Y%m%dT%H%M%S%L%Z", t)
	return strings.Replace(timestamp, ".", "", 1)
}

// TimestampWithDelimiter returns a timestamp string for the specified
// time in a standart format using the specified delimeter to separate
// the different components of the date (year, month, day and time),
// e.g. 2017/01/01/123401CET.
func TimestampWithDelimiter(t time.Time, d string) string {
	timestamp := strftime.Format("%Y-%m-%d-%H%M%S%L%Z", t)
	timestamp = strings.Replace(timestamp, ".", "", 1)
	return strings.Replace(timestamp, "-", d, -1)
}
