package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	if target == 'z' || target < letters[0] {
		return letters[0]
	}

	for i, b := range letters {
		if target < b {
			return letters[i]
		}
	}

	return letters[0]
}
