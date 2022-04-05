package leetcode

func countPrimeSetBits(left, right int) int {
	ans, primes := 0, map[int]bool{2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true}
	for i := left; i <= right; i++ {
		cur := 0
		for j := i; j > 0; j -= j & (-j) {
			cur++
		}
		if primes[cur] {
			ans++
		}
	}
	return ans
}
