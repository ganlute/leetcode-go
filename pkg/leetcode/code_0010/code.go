package code_0010

// 这道题写的什么玩意儿～
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) == 0 {
		if len(p)%2 == 1 {
			return false
		}
		for i := 1; i < len(p); i = i + 2 {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}
	if len(p) == 0 {
		return false
	}

	m, n := len(s), len(p)
	dp := make([][]int, m+1) // 0代表未知，1代表false，2代表true
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 初始化
	dp[0][0] = 2
	dp[0][1] = 1
	dp[1][0] = 1
	if s[0] == p[0] || p[0] == '.' {
		dp[1][1] = 2
	} else {
		dp[1][1] = 1
	}
	for j := 2; j <= n; j++ {
		if p[j-1] != '*' {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 2; i <= m; i++ {
		dp[i][0] = 1
		dp[i][1] = 1
	}

	// dp[i][j]
	for i := 1; i <= m; i++ {
		for j := 2; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					if dp[i-1][j] == 2 || dp[i][j-1] == 2 || dp[i][j-2] == 2 {
						dp[i][j] = 2
					} else {
						dp[i][j] = 1
					}
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				dp[i][j] = 1
			}
		}
	}
	return dp[m][n] == 2
}
