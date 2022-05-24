package leetcode

import (
	"sort"
)

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int, len(words))
	for _, w := range words {
		m[w]++
	}

	pl := make(pairList, 0, len(m))
	for w, c := range m {
		pl = append(pl, &pair{
			key:   w,
			value: c,
		})
	}

	sort.Sort(pl)

	ans := make([]string, k)
	for i := 0; i < k; i++ {
		ans[i] = pl[i].key
	}

	return ans
}

type pair struct {
	key   string
	value int
}

type pairList []*pair

func (p pairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pairList) Len() int      { return len(p) }
func (p pairList) Less(i, j int) bool {
	if p[i].value == p[j].value {
		return p[i].key < p[j].key
	}

	return p[i].value > p[j].value
}
