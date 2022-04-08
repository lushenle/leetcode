package leetcode

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	if sx > tx || sy > ty {
		return false
	}

	for sx < tx && sy < ty {
		if tx < ty {
			ty %= tx
		} else {
			tx %= ty
		}
	}

	return sx == tx && (ty-sy)%sx == 0 || sy == ty && (tx-sx)%sy == 0
}
