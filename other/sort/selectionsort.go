package sort

// findSmallest find the smallest element in the array and return its index
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIdx := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIdx = i
		}
	}

	return smallestIdx
}

func selectionSort(arr []int) []int {
	var result []int
	count := len(arr)
	for i := 0; i < count; i++ {
		smallestIdx := findSmallest(arr)
		result = append(result, arr[smallestIdx])
		arr = append(arr[:smallestIdx], arr[smallestIdx+1:]...)
	}

	return result
}
