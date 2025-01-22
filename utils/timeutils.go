package utils

import "time"

func EpochMillisToTime(epochMillis int64) string {
	t := time.Unix(0, epochMillis*int64(time.Millisecond))
	return t.Format("2006-01-02 15:04")
}
