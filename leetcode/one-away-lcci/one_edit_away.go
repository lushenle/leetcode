package leetcode

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway(second, first)
	}

	if m-n > 1 {
		return false
	}

	if m == n {
		diff := false
		for i := 0; i < n; i++ {
			if first[i] != second[i] {
				if diff {
					return false
				}
				diff = true
			}
		}
		return true
	}

	i, j, cnt := 0, 0, 0
	for i < m && j < n {
		if first[i] == second[j] {
			i, j = i+1, j+1
		} else {
			i, cnt = i+1, cnt+1
		}
	}
	return cnt <= 1
}
