package main

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
		// dp[i][0]=false
	}
	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '?' {
				if i == 0 {
					// dp[i][j]=false
				} else {
					dp[i][j] = dp[i-1][j-1]
				}
			} else if p[j-1] == '*' {
				if i == 0 {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i][j-1] || dp[i-1][j]
				}
			} else {
				if i > 0 {
					dp[i][j] = s[i-1] == p[j-1] && dp[i-1][j-1]
				} // else false
			}
		}
	}
	return dp[len(s)][len(p)]
}
func allStar(p string) bool {
	for _, c := range []byte(p) {
		if c != '*' {
			return false
		}
	}
	return true
}
