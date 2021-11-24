package leetcode

// Ref:
//  https://en.wikipedia.org/wiki/Fast_inverse_square_root
//  https://www.zhihu.com/question/20690553

//func mySqrt(x int) int {
//	if x == 0 || x == 1 {
//		return x
//	}
//
//	for i := 1; i <= x; i++ {
//		if x-i*i < 0 {
//			return i - 1
//		}
//	}
//
//	return -1
//}

//func mySqrt(x int) int {
//	if x == 0 || x == 1 {
//		return x
//	}
//
//	result := -1
//	left, right := 0, x
//
//	for left <= right {
//		mid := left + (right-left)/2
//		if mid*mid <= x {
//			result = mid
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//	return result
//}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}

	return r
}
