package leetcode

func isPowerOfTwo(n int) bool {
	//return n > 0 && n&(n-1) == 0
	//return n > 0 && n&-n == n
	return n > 0 && 1073741824%n == 0
}
