package shuffle

import (
	"math/rand"
	"time"
)

func Slice[T any](a []T) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return a
}
