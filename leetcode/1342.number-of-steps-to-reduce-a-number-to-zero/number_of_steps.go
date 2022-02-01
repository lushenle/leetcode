package leetcode

import "math"

func numberOfSteps(num int) int {
	if num == 0 {
		return num
	}
	return bitCount(uint32(num)) + int(math.Floor(math.Log2(float64(num))))
}

func bitCount(n uint32) int {
	n = n - ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n + (n >> 4)) & 0x0F0F0F0F
	return int((n * 0x01010101) >> 24)
}
