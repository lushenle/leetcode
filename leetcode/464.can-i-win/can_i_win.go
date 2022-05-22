package leetcode

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	// sum 是所有能选数之和
	sum := (1 + maxChoosableInteger) * maxChoosableInteger / 2
	// 这种情况下，双方都不能赢
	if sum < desiredTotal {
		return false
	}

	// bit[i] 等于，仅在第 i 个 bit 位为 1 的整数值
	bit := make([]int, maxChoosableInteger+1)
	for i := maxChoosableInteger; i > 0; i-- {
		bit[i] = 1 << uint8(i)
	}

	// dp[42] == true 表示
	// 面临 1,3,5 被选择的人，有机会赢
	dp := make(map[int]bool, maxChoosableInteger*maxChoosableInteger)

	return dfs(0, desiredTotal, maxChoosableInteger, dp, bit)
}

func dfs(usedBit, remains, max int, dp map[int]bool, bit []int) bool {
	// 已经计算过 usedBit 这种状态了，直接返回结果
	if res, ok := dp[usedBit]; ok {
		return res
	}

	if remains <= max {
		for i := max; i >= remains; i-- {
			if usedBit&bit[i] == 0 {
				dp[usedBit] = true
				return true
			}
		}
	}

	for i := max; i > 0; i-- {
		if usedBit&bit[i] > 0 {
			// i 已经被使用过了
			continue
		}

		old := usedBit
		// 暂时选择 i ，往 usedBit 中添加 i
		usedBit |= bit[i]
		isOpponentWin := dfs(usedBit, remains-i, max, dp, bit)
		// 在 usedBit 中删除 i
		usedBit = old
		// usedBit &= (^bit[i])

		if !isOpponentWin { // 对手输了
			// 所以，是我赢了
			// 请注意，下面的 usedBit 是删除了 i 的
			// 此时的 dp[usedBit] = true 表示
			// 我在面临 usedBit 这种状态下能赢，因为可以选择 i
			dp[usedBit] = true
			return true
		}
	}

	dp[usedBit] = false
	return false
}
