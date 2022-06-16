package leetcode

func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	ans := 0

	for _, num := range nums {
		m[num]++
	}

	if k == 0 {
		for _, count := range m {
			if count > 1 {
				ans++
			}
		}
		return ans
	}

	for n := range m {
		if m[n-k] > 0 {
			ans++
		}
	}

	return ans
}
