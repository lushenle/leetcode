package leetcode

import "container/heap"

//func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
//	var result [][]int
//	for i := 0; i < len(nums1) && i < k; i++ {
//		for j := 0; j < len(nums2) && j < k; j++ {
//			result = append(result, []int{nums1[i], nums2[j]})
//		}
//	}
//
//	sort.Slice(result, func(i, j int) bool {
//		return result[i][0]+result[i][1] < result[j][0]+result[j][1]
//	})
//	if len(result) >= k {
//		return result[:k]
//	}
//	return result
//}

// minHeap need to be implementation: Len, Less, Swap, Push, Pop methods
type minHeap [][]int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// kSmallestPairs return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.
// !!!
//
// - nums1: [1,1,2]
// - nums2: [1,2,3]
// - k:3
// return: [[1,1],[1,1],[2,1]] is ok
// return: [[1,1],[1,1],[1,2]] is not ok !!!
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result [][]int
	h := &minHeap{}

	// Constraints:
	// 1 <= nums1.length, nums2.length <= 105
	// -109 <= nums1[i], nums2[i] <= 109
	// nums1 and nums2 both are sorted in ascending order
	// 1 <= k <= 1000
	//if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
	//	return result
	//}

	if len(nums1)*len(nums2) < k {
		k = len(nums1) * len(nums2)
	}

	heap.Init(h)

	for _, num := range nums1 {
		heap.Push(h, []int{num, nums2[0], 0})
	}
	for len(result) < k {
		min := heap.Pop(h).([]int)
		result = append(result, min[:2])

		if min[2] < len(nums2)-1 {
			heap.Push(h, []int{min[0], nums2[min[2]+1], min[2] + 1})
		}
	}
	return result
}
