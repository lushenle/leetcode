package leetcode

func strStr(haystack string, needle string) int {
	//if len(needle) == 0 {
	//	return 0
	//}
	//
	//if !strings.Contains(haystack, needle) || len(haystack) == 0 {
	//	return -1
	//}

	// solution 1
	//return strings.Index(haystack, needle)

	// solution 2
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}

			if i+j == len(haystack) {
				return -1
			}

			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
