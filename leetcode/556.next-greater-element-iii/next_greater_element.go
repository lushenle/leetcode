package leetcode

import (
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	s := []byte(strconv.Itoa(n))
	i := len(s) - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i == -1 {
		return -1
	}
	j := len(s) - 1
	for j >= 0 && s[i] >= s[j] {
		j--
	}

	s[i], s[j] = s[j], s[i]
	temp := s[i+1:]
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})

	ans, _ := strconv.Atoi(string(append(s[:i+1], temp...)))
	if ans > math.MaxInt32 {
		return -1
	}

	return ans
}
