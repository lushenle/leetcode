package leetcode

import (
	"strings"
)

func countValidWords(sentence string) int {
	ans := 0

	for _, s := range strings.Split(sentence, " ") {
		if isStr(s) {
			ans++
		}
	}

	return ans
}

func isStr(s string) bool {
	l := len(s)

	if l == 0 {
		return false
	}

	if s[0] == '-' || s[l-1] == '-' {
		return false
	}

	if !('a' <= s[0] || s[0] <= 'z') {
		return false
	}

	var count int
	for i := 0; i < l; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return false
		}

		if i < l-1 && (s[i] == '.' || s[i] == '!' || s[i] == ',') {
			return false
		}

		if s[i] == '-' {
			count++
			if count == 2 {
				return false
			}

			if i-1 < 0 || !(s[i-1] >= 'a' && s[i-1] <= 'z') || i+1 >= l || !(s[i+1] >= 'a' && s[i+1] <= 'z') {
				return false
			}
		}
	}

	return true
}

//func countValidWords1(sentence string) int {
//	ans := 0
//
//	re := regexp.MustCompile(`^[a-z]*([a-z]-[a-z])?[a-z]*[!.,]?$`)
//	for _, s := range strings.Fields(sentence) {
//		if re.MatchString(s) {
//			ans++
//		}
//	}
//
//	return ans
//}
