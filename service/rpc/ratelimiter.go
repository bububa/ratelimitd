package rpc

import (
	"time"

	"go.uber.org/ratelimit"

	"github.com/bububa/ratelimitd/pb"
)

type Limiter struct {
	conf   *pb.Limiter
	bucket ratelimit.Limiter
}

func NewLimiter(conf *pb.Limiter) *Limiter {
	return &Limiter{
		conf:   conf,
		bucket: ratelimit.New(int(conf.GetRate()), ratelimit.Per(time.Duration(conf.GetInterval()))),
	}
}

func (l *Limiter) Take() time.Time {
	return l.bucket.Take()
}

func (l *Limiter) Config() *pb.Limiter {
	return l.conf
}
