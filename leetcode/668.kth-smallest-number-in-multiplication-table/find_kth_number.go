package leetcode

import "sort"

func findKthNumber(m int, n int, k int) int {
	// 给定一个数x，求矩阵中 <=x 的数的个数，利用右上角二分
	lessEqual := func(x int) int {
		i := 0
		j := n - 1
		count := 0
		for i < m && j >= 0 {
			num := (i + 1) * (j + 1)
			if num <= x {
				count += j + 1
				i++
			} else {
				j--
			}
		}
		return count
	}

	// 二分找到 [0, m*n] 中，第一个满足条件的x
	return sort.Search(m*n+1, func(x int) bool {
		return lessEqual(x) >= k
	})
}
