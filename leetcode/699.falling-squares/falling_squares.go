package leetcode

import (
	"container/heap"
	"runtime/debug"
)

func fallingSquares(positions [][]int) []int {
	ans := make([]int, 0, len(positions))
	pq := make(PQ, 0, len(positions))

	for i := 0; i < len(positions); i++ {
		sp := &segment{
			left:   positions[i][0],
			right:  positions[i][0] + positions[i][1],
			height: positions[i][1],
		}

		height := 0
		removes := make([]*segment, 0, len(pq))

		for j := 0; j < len(pq); j++ {
			if pq[j].right <= sp.left || sp.right <= pq[j].left {
				continue
			}

			height = max(height, pq[j].height)

			if sp.left <= pq[j].left && pq[j].right <= sp.right {
				removes = append(removes, pq[j])
			}

			if pq[j].left < sp.left && sp.right < pq[j].right {
				heap.Push(&pq, &segment{
					left:   sp.right,
					right:  pq[j].right,
					height: pq[j].height,
				})
				pq[j].right = sp.left
				break
			}

			if pq[j].left < sp.left && sp.left < pq[j].right && pq[j].right <= sp.right {
				pq[j].right = sp.left
			}

			if sp.left <= pq[j].left && pq[j].left < sp.right && sp.right < pq[j].right {
				pq[j].left = sp.right
			}
		}

		for j := 0; j < len(removes); j++ {
			heap.Remove(&pq, removes[j].index)
		}

		sp.height += height
		heap.Push(&pq, sp)
		ans = append(ans, pq[0].height)
	}

	return ans
}

type segment struct {
	left, right int
	height      int
	index       int
}

type PQ []*segment

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].height > pq[j].height }

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	temp := x.(*segment)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() { debug.SetGCPercent(-1) }
