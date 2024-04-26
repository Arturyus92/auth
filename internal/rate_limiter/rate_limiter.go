package rate_limiter

import (
	"context"
	"time"
)

// TokenBucketLimiter ...
type TokenBucketLimiter struct {
	tokenBucketCh chan struct{}
}

// NewTokenBucketLimiter ...
func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokenBucketCh: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokenBucketCh <- struct{}{}
	}

	repleninshmentInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicReplenishment(ctx, time.Duration(repleninshmentInterval))

	return limiter
}

func (l *TokenBucketLimiter) startPeriodicReplenishment(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			l.tokenBucketCh <- struct{}{}
		}
	}
}

// Allow ...
func (l *TokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokenBucketCh:
		return true
	default:
		return false
	}
}
