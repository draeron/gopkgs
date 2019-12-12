package chrono

import (
	"sync"
	"time"
)

type StopWatch struct {
	start   time.Time
	elapsed time.Duration
	running bool

	mutex sync.RWMutex
}

func NewStopWatch(start bool) *StopWatch {
	c := &StopWatch{}
	if start {
		c.Start()
	}
	return c
}

func (c *StopWatch) Start() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.start = time.Now()
	c.running = true
}

func (c *StopWatch) Stop() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.running {
		c.elapsed += time.Since(c.start)
		c.running = false
	}
}

func (c *StopWatch) Add(duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.elapsed += duration
}

func (c *StopWatch) ResetTo(duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.elapsed = duration
	c.start = time.Now()
}

func (c *StopWatch) IsRunning() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.running
}

func (c *StopWatch) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.elapsed = 0
	c.running = false
}

func (c *StopWatch) Elapsed() time.Duration {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	elap := c.elapsed
	if c.running {
		elap += time.Since(c.start)
	}
	return elap
}
