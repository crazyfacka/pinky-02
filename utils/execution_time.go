package utils

import (
	"time"
)

func ExecutionTime(fn func()) time.Duration {
	t1 := time.Now()
	fn()
	t2 := time.Now()
	return t2.Sub(t1)
}
