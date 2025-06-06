package limiters

import (
	"github.com/luk3skyw4lker/rate-limiter/config"
	"github.com/luk3skyw4lker/rate-limiter/store"
)

type FixedWindowLimiter struct {
	maxCount int
	store    store.Store
}

func NewFixedWindowLimiter(limiterConfig config.Config) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		maxCount: limiterConfig.MaxCount,
		store:    limiterConfig.Store,
	}
}

func (l *FixedWindowLimiter) Allow(clientId string) bool {
	currentCount, err := l.store.GetCurrentCount(clientId)
	if err != nil {
		return false
	}

	if currentCount >= l.maxCount {
		return false
	}

	newCount, err := l.store.IncrementCount(clientId)
	if err != nil {
		return false
	}

	return newCount <= l.maxCount
}

func (l *FixedWindowLimiter) Reset(clientId string) error {
	return l.store.ResetCount(clientId)
}
