package random

import (
	"math/rand"
	"time"
)

// MinMaxInt returns random value between min inclusive and max exclusive [min,max)
func MinMaxInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}
