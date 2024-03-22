package rand

import (
	"math/rand"
	"time"
)

func Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

const (
	_MIN = 10_000
	_MAX = 99_999
)

func IntStandard() int {
	return Int(_MIN, _MAX)
}
