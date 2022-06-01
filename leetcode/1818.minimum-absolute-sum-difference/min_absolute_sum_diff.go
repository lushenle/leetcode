package leetcode

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	diff, maxDiff := 0, 0
	for idx, n2 := range nums2 {
		d := abs(nums1[idx] - n2)
		diff += d
		if maxDiff < d {
			t := 100001
			for _, n1 := range nums1 {
				maxDiff = max(maxDiff, d-min(t, abs(n1-n2)))
			}
		}
	}

	return (diff - maxDiff) % (1e9 + 7)
}

func abs(a int) int        { return (a ^ (a >> 63)) - (a >> 63) }
func max(a int, b int) int { return (a + b + abs(a-b)) / 2 }
func min(a, b int) int     { return a + (b-a)>>31&(b-a) }
