package input

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomRerolls() []int {
	rerolls := []int{}
	for i := 0; i <= 4; i += 1 {
		rand := rand.Float32()
		if rand >= 0.5 {
			rerolls = append(rerolls, i)
		}
	}
	return rerolls
}
