package studygov2

import "time"

func DateFormat(t time.Time) int {
	return t.Hour()
}