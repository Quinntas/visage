package timeUtils

import "time"

func TimeDurationToTime(duration time.Duration) time.Time {
	return time.Unix(int64(duration/time.Second), 0)
}
