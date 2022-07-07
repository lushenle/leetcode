package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1 == 0 && n2 == 0 && n3 == 0 {
		return true
	}

	if n1+n2 != n3 {
		return false
	}

	dp := make([][][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([][]bool, n2+1)
		for j := 0; j <= n2; j++ {
			dp[i][j] = make([]bool, n3+1)
		}
	}
	dp[0][0][0] = true

	for i := 1; i <= n1; i++ {
		dp[i][0][i] = dp[i-1][0][i-1] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= n2; j++ {
		dp[0][j][j] = dp[0][j-1][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			dp[i][j][i+j] = (dp[i-1][j][i+j-1] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1][i+j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[n1][n2][n3]
}
