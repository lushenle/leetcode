package leetcode

func platesBetweenCandles(s string, queries [][]int) []int {
	// magic[i]: [candle index, left valid plates]
	magic := make([][]int, 0)
	ans := make([]int, len(queries))
	
	for i, v := range s {
		if v == '|' {
			magic = append(magic, []int{i, 0})
		}
	}

	if len(magic) <= 1 {
		// no candle exists or only one candle, all query ans must be 0
		return ans
	}

	for i := 1; i < len(magic); i++ {
		magic[i][1] = magic[i][0] - magic[i-1][0] - 1 + magic[i-1][1]
	}

	for i, q := range queries {
		// find GE for indexL
		indexL := binarySearchGE(magic, q[0])

		if indexL == -1 {
			continue
		}
		// find LE for indexR
		indexR := binarySearchGE(magic, q[1])
		if indexR == 0 {
			// GE index is 0, query ans must be 0
			continue
		} else if indexR == -1 {
			// GE index not exist, LE index must be len(magic)-1
			indexR = len(magic) - 1
		} else if magic[indexR][0] != q[1] {
			// not equal, LE must be GE-1
			indexR--
		}

		if magic[indexR][1] > magic[indexL][1] {
			ans[i] = magic[indexR][1] - magic[indexL][1]
		}
	}

	return ans
}

// find first GE index of candle, -1 means not found
func binarySearchGE(magic [][]int, index int) int {
	l, r := 0, len(magic)-1
	for l <= r {
		m := (l + r) >> 1
		if magic[m][0] >= index {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if r+1 > len(magic)-1 {
		return -1
	}

	return r + 1
}
