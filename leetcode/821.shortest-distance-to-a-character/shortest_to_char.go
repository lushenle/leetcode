package leetcode

/* 先找到所有的 c 的位置
* 再遍历 s, 记录每个位置到 c 的距离
* 最后返回最短距离
 */

func shortestToChar(s string, c byte) []int {
	var ans []int
	var pos []int

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			pos = append(pos, i)
		}
	}

	for i := 0; i < len(s); i++ {
		min := len(s)
		for _, v := range pos {
			if abs(i-v) < min {
				min = abs(i - v)
			}
		}
		ans = append(ans, min)
	}

	return ans
}

func abs(a int) int {
	b := a >> 63
	return (a ^ b) - b
}
