package leetcode

func luckyNumbers(matrix [][]int) []int {
	var ans []int
	t, r := make([]int, len(matrix[0])), make([]int, len(matrix[0]))

	for _, val := range matrix {
		m, k := val[0], 0
		for j := 0; j < len(matrix[0]); j++ {
			if val[j] < m {
				m = val[j]
				k = j
			}
			if t[j] < val[j] {
				t[j] = val[j]
			}
		}

		if t[k] == m {
			r[k] = m
		}
	}

	for k, v := range r {
		if v > 0 && v == t[k] {
			ans = append(ans, v)
		}
	}

	return ans
}
