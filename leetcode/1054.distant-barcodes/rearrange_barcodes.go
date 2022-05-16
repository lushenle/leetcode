package leetcode

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	m := map[int]int{}

	for _, v := range barcodes {
		m[v]++
	}

	content := make([][]int, 0, len(m))
	for k, v := range m {
		content = append(content, []int{k, v})
	}

	sort.Slice(content, func(i, j int) bool {
		return content[i][1] > content[j][1]
	})

	idx := 0
	for i := 0; i < len(barcodes); i += 2 {
		if content[idx][1] == 0 {
			idx++
		}
		content[idx][1]--
		barcodes[i] = content[idx][0]
	}

	for i := 1; i < len(barcodes); i += 2 {
		if content[idx][1] == 0 {
			idx++
		}
		content[idx][1]--
		barcodes[i] = content[idx][0]
	}

	return barcodes
}
