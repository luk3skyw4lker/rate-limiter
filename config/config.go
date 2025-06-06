package config

import "github.com/luk3skyw4lker/rate-limiter/store"

type Config struct {
	MaxCount int
	Store    store.Store
}
