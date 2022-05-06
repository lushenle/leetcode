package leetcode

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	for len(rc.queue) > 0 && rc.queue[0] < t-3000 {
		rc.queue = rc.queue[1:]
	}

	rc.queue = append(rc.queue, t)
	return len(rc.queue)
}
