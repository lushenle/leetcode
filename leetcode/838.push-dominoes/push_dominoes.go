package leetcode

func push(d []byte) {
	first, last := 0, len(d)-1
	switch d[first] {
	case '.', 'L':
		if d[last] == 'L' {
			for ; first < last; first++ {
				d[first] = 'L'
			}
		}
	case 'R':
		if d[last] == '.' || d[last] == 'R' {
			for ; first <= last; first++ {
				d[first] = 'R'
			}
		} else if d[last] == 'L' {
			for first < last {
				d[first] = 'R'
				d[last] = 'L'
				first++
				last--
			}
		}
	}
}

func pushDominoes(dominoes string) string {
	d := []byte(dominoes)
	for i := 0; i < len(d); {
		j := i + 1
		for j < len(d)-1 && d[j] == '.' {
			j++
		}
		push(d[i : j+1])
		i = j
	}
	return string(d)
}
