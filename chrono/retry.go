package chrono

import "time"

func DoRetry(delay time.Duration, retry int, do func() error) {
	tick := time.NewTicker(delay)
	defer tick.Stop()

	for i := 0; i <= retry; i++ {
		err := do()
		if err == nil {
			return
		}
		<-tick.C
	}
}
