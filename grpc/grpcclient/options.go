package grpcclient

import "time"

type Config struct {
	timeout time.Duration
}

type Option func(*Config)

func Timeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.timeout = timeout
	}
}
