package utils

import (
	"math/rand"
	"time"
)

// GetRandomArray ...
func GetRandomArray(length int) []int {
	rand.Seed(time.Now().Unix())
	var src, dest []int
	for i := 0; i < length; i++ {
		src = append(src, i+1)
	}
	for {
		if len(src) == 0 {
			break
		}
		i := rand.Intn(len(src))
		dest = append(dest, src[i])
		src = append(src[:i], src[i+1:]...)
	}
	return dest
}
