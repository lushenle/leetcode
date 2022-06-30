package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := 0, 0
	n := len(intervals)
	for ; left < len(intervals) && newInterval[0] > intervals[left][1]; left++ {
	}

	for right = left; right < len(intervals) && newInterval[1] >= intervals[right][0]; right++ {
	}

	if left == right {
		intervals = append(intervals, []int{0, 0})
		copy(intervals[right+1:], intervals[right:])
		intervals[right] = newInterval
	} else {
		if newInterval[0] < intervals[left][0] {
			intervals[left][0] = newInterval[0]
		}
		if newInterval[1] > intervals[right-1][1] {
			intervals[left][1] = newInterval[1]
		} else {
			intervals[left][1] = intervals[right-1][1]
		}
		copy(intervals[left+1:], intervals[right:])
	}

	return intervals[:left+n-right+1]
}
