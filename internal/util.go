package internal

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func generate() []int {
	var nums []int
	elementsNum := randomInt(1, 32768)

	for i := 0; i <= elementsNum; i++ {
		nums = append(nums, rand.Int())
	}

	return nums
}
