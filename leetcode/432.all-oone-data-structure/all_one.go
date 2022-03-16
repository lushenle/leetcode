package leetcode

import "container/heap"

type entry struct {
	key      string
	value    int
	maxIndex int
	minIndex int
}

type maxPQ []*entry

func (pq maxPQ) Len() int           { return len(pq) }
func (pq maxPQ) Less(i, j int) bool { return pq[i].value > pq[j].value }
func (pq maxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].maxIndex, pq[j].maxIndex = i, j
}

func (pq *maxPQ) Push(x interface{}) {
	n := len(*pq)
	e := x.(*entry)
	e.maxIndex = n
	*pq = append(*pq, e)
}

func (pq *maxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	e := old[n-1]
	e.maxIndex = -1
	*pq = old[0 : n-1]
	return e
}

type minPQ []*entry

func (pq minPQ) Len() int           { return len(pq) }
func (pq minPQ) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].minIndex, pq[j].minIndex = i, j
}

func (pq *minPQ) Push(x interface{}) {
	n := len(*pq)
	e := x.(*entry)
	e.minIndex = n
	*pq = append(*pq, e)
}

func (pq *minPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	e := old[n-1]
	e.minIndex = -1
	*pq = old[0 : n-1]
	return e
}

type AllOne struct {
	m   map[string]*entry
	max maxPQ
	min minPQ
}

func Constructor() AllOne {
	return AllOne{
		m:   make(map[string]*entry),
		max: maxPQ{},
		min: minPQ{},
	}
}

func (a *AllOne) Inc(key string) {
	if ep, ok := a.m[key]; ok {
		ep.value++
		heap.Fix(&a.max, ep.maxIndex)
		heap.Fix(&a.min, ep.maxIndex)
	} else {
		ep = &entry{key: key, value: 1}
		a.m[key] = ep
		heap.Push(&a.max, ep)
		heap.Push(&a.min, ep)
	}
}

func (a *AllOne) GetMaxKey() string {
	if len(a.max) == 0 {
		return ""
	}

	return a.max[0].key
}

func (a *AllOne) Dec(key string) {
	if ep, ok := a.m[key]; !ok {
		return
	} else if ep.value == 1 {
		heap.Remove(&a.max, ep.maxIndex)
		heap.Remove(&a.min, ep.minIndex)
	} else {
		ep.value--
		heap.Fix(&a.max, ep.maxIndex)
		heap.Fix(&a.min, ep.minIndex)
	}
}

func (a *AllOne) GetMinKey() string {
	if len(a.max) == 0 {
		return ""
	}

	return a.min[0].key
}
