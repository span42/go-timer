package timer

import (
	"time"
)

const delay = 10000

type Timer struct {
	timer  *time.Timer
	lpFunc func(*Timer, ...interface{})
	exitCh chan bool
}

func NewTimer(lpFunc func(*Timer, ...interface{}), lpArg ...interface{}) *Timer {
	nt := new(Timer)
	nt.timer = time.NewTimer(time.Duration(delay) * time.Millisecond)
	nt.timer.Stop()
	nt.lpFunc = lpFunc
	nt.exitCh = make(chan bool)
	go func(*Timer) {
		for {
			select {
			case <-nt.timer.C:
				if nt.lpFunc != nil {
					nt.lpFunc(nt, lpArg...)
				}
			case <-nt.exitCh:
				nt.timer.Stop()
				return
			}
		}
	}(nt)
	return nt
}

func (t *Timer) Start(msec int) {
	t.timer.Reset(time.Duration(msec) * time.Millisecond)
}
func (t *Timer) Stop() {
	t.timer.Stop()
}
func (t *Timer) Destory() {
	t.exitCh <- true
}
