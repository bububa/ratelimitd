package pb

import "time"

func NewLimiter(name string, rate int, interval time.Duration) *Limiter {
	ret := new(Limiter)
	ret.Name = name
	ret.Rate = int32(rate)
	ret.Interval = interval.Nanoseconds()
	return ret
}

func (l *Limiter) IsValid() bool {
	return l.GetRate() > 0 && l.GetInterval() > 0
}
