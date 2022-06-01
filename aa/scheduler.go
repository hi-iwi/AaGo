package aa

import (
	"github.com/RussellLuo/timingwheel"
	"time"
)

type EveryScheduler struct {
	Interval time.Duration
}

func (s *EveryScheduler) Next(prev time.Time) time.Time {
	return prev.Add(s.Interval)
}

func (app *Aa) ScheduleFunc(s timingwheel.Scheduler, f func()) (t *timingwheel.Timer) {
	return app.wheelTimer.ScheduleFunc(s, f)
}

func (app *Aa) AfterFunc(d time.Duration, f func()) *timingwheel.Timer {
	return app.wheelTimer.AfterFunc(d, f)
}

func (app *Aa) Schedule(d time.Duration, f func()) (t *timingwheel.Timer) {
	return app.wheelTimer.ScheduleFunc(&EveryScheduler{d}, f)
}
