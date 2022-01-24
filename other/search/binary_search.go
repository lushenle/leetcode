package search

// binarySearch return the index of the target
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for {
		if low <= high {
			mid := (low + high) / 2
			guess := arr[mid]
			if guess == target {
				return mid
			}
			if guess > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
}
