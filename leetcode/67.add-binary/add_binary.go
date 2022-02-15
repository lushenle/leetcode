package leetcode

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}

	if len(b) > len(a) {
		a, b = b, a
	}

	ans := make([]string, len(a)+1)
	i, j, k, c := len(a)-1, len(b)-1, len(a), 0
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		ans[k] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2
		i--
		j--
		k--
	}

	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		ans[k] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		k--
	}

	if c > 0 {
		ans[k] = strconv.Itoa(c)
	}

	return strings.Join(ans, "")
}
