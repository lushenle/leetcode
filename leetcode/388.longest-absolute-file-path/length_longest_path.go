package leetcode

import "strings"

func lengthLongestPath(input string) int {
	if !hasFile(input) {
		return 0
	}

	return llp("\n"+input) - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func nextCR(path string, idx int) int {
	var i int
	for i = idx + 1; i < len(path); i++ {
		if path[i:i+1] == "\n" && (i+1 == len(path) || path[i+1:i+2] != "\t") {
			break
		}
	}

	return i
}

func hasSub(path string) bool {
	return strings.Contains(path, "\n")
}

func hasFile(path string) bool {
	return strings.Contains(path, ".")
}

func llp(path string) int {
	if !hasFile(path) {
		return 0
	}

	if !hasSub(path) {
		return len(path)
	}

	var i = nextCR(path, -1)
	dirLen := i

	var j, maxSub int
	for i < len(path) {
		j = nextCR(path, i)
		maxSub = max(maxSub, llp(strings.Replace(path[i+1:j], "\n\t", "\n", -1)))
		i = j
	}

	return dirLen + 1 + maxSub
}
