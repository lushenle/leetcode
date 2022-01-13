package subset

// isSubset determine if the non-empty set s2 is a subset of s1
func isSubset(s1, s2 []int) bool {
	e := make(map[int]int)

	for i, v := range s1 {
		e[v] = i
	}

	for _, v2 := range s2 {
		if _, ok := e[v2]; !ok {
			return false
		}
	}

	return true
}
