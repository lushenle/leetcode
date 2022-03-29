package leetcode

func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	for left, right, cntsT, cntsF := 0, 0, 0, 0; right < len(answerKey); right++ {
		cntsT += compareStr(answerKey[right], 'T')
		cntsF += compareStr(answerKey[right], 'F')

		for cntsT > k && cntsF > k {
			cntsT -= compareStr(answerKey[left], 'T')
			cntsF -= compareStr(answerKey[left], 'F')
			left++
		}

		if length := right - left + 1; length > ans {
			ans = length
		}
	}

	return ans
}

func compareStr(a, b byte) int {
	if a == b {
		return 1
	}

	return 0
}
