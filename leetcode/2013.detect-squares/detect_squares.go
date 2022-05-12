package leetcode

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

type DetectSquares struct {
	// 横坐标x, 纵坐标y, (x,y) 位置上点的个数
	Point map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		Point: map[int]map[int]int{},
	}
}

func (ds *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]

	if _, ok1 := ds.Point[x]; ok1 {
		ds.Point[x][y] += 1
	} else {
		ds.Point[x] = map[int]int{}
		ds.Point[x][y] = 1
	}
}

func (ds *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	count := 0

	// 如果哈希表 Point 中存在横坐标为 x 的点集合 ySet
	if ySet, ok := ds.Point[x]; ok {
		// 遍历 ySet, 对每个不同于 y 的纵坐标 y1, 计算正方形边长
		for y1 := range ySet {
			if y1 == y {
				continue
			}
			var distance int
			if y1 > y {
				distance = y1 - y
			} else {
				distance = y - y1
			}

			// 找到左边的点，判断是否能构成正方形
			if yLeft, ok1 := ds.Point[x-distance]; ok1 {
				y2, ok2 := yLeft[y]  // 传入点的左边的个数
				y3, ok3 := yLeft[y1] // 与传入点的同一个坐标轴的左边的个数
				if ok2 && ok3 {
					count += ySet[y1] * y2 * y3
				}
			}

			// 找到右边的点，判断是否能构成正方形
			if yRight, ok1 := ds.Point[x+distance]; ok1 {
				y2, ok2 := yRight[y]  // 传入点的右边的个数
				y3, ok3 := yRight[y1] // 与传入点的同一个坐标轴的右边的个数
				if ok2 && ok3 {
					count += ySet[y1] * y2 * y3
				}
			}
		}
	}

	return count
}
