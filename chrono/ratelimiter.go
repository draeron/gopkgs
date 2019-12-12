package chrono

import (
	"context"
	"github.com/TheCodeTeam/goodbye"
	"go.uber.org/atomic"
	"os"
	"sync"
	"time"
)

type DelayedExec struct {
	fct   func()
	limit time.Duration
	wait  *sync.Cond
	drop  bool
	done  atomic.Bool
}

func NewDelayedExec(duration time.Duration, waitIfTooFast bool) *DelayedExec {
	d := &DelayedExec{
		limit: duration,
		drop:  !waitIfTooFast,
		wait:  sync.NewCond(&sync.Mutex{}),
	}

	goodbye.Register(func(ctx context.Context, s os.Signal) {
		d.done.Store(true)
	})

	go d.run()
	return d
}

func (d *DelayedExec) Exec(_fct func()) {
	d.wait.L.Lock()
	//if d.fct != nil {
	//	log.Debug("overriding previous")
	//}
	d.fct = _fct
	d.wait.L.Unlock()
	d.wait.Broadcast()
}

func (d *DelayedExec) run() {

	last := time.Now()

	for !d.done.Load() {

		d.wait.L.Lock()
		if d.fct == nil {
			d.wait.Wait()
			//log.Debug("woke up")
		}
		d.wait.L.Unlock()

		diff := time.Now().Sub(last)
		if diff < d.limit {
			//log.Sugar().Debug("too close waiting ", diff.String())

			if !d.drop {
				<-time.After(diff)
			} else {
				d.wait.L.Lock()
				d.fct = nil
				d.wait.L.Unlock()
				continue
			}
		}

		d.wait.L.Lock()
		d.fct()
		d.fct = nil
		last = time.Now()
		d.wait.L.Unlock()
	}
}
