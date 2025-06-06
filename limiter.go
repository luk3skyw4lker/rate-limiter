package ratelimiter

import (
	"github.com/luk3skyw4lker/rate-limiter/config"
	"github.com/luk3skyw4lker/rate-limiter/store"
)

type RateLimiter struct {
	maxCount int
	store    store.Store
}

func NewRateLimiter(limiterConfig config.Config) *RateLimiter {
	return &RateLimiter{
		maxCount: limiterConfig.MaxCount,
		store:    limiterConfig.Store,
	}
}

func (l *RateLimiter) Allow(clientId string) bool {
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
