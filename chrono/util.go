package chrono

import (
	"fmt"
	"time"
)

func FmtDuration(d time.Duration) string {
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond
	return fmt.Sprintf("%02d:%02d:%02d", m, s, ms/10)
}

func SplitDuration(d time.Duration) [3]uint {
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	return [3]uint{uint(m), uint(s / 10), uint(s % 10)}
}
