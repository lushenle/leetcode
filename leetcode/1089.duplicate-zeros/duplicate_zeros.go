package leetcode

func duplicateZeros(arr []int) {
	var shifted []int

	for idx, n := range arr {
		if n == 0 {
			shifted = append(shifted, 0)
			shifted = append(shifted, 0)
		} else {
			shifted = append(shifted, arr[idx])
		}

		if len(shifted) == len(arr) {
			break
		}
	}

	copy(arr, shifted)
}
