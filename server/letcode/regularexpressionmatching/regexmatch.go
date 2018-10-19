package regexmatch

// IsMatch is a function that to judge if a string s is match regex p
func IsMatch(s string, p string) bool {
	t := 0
	for i := 0; i < len(p); i++ {
		if p[i:i+1] == "*" {
			continue
		} else {
			t = i
			break
		}
	}
	p = p[t:len(p)]
	var dp [][]bool
	for i := 0; i < len(s)+1; i++ {
		tmp := make([]bool, len(p)+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for i := 0; i < len(p); i++ {
		if p[i:i+1] == "*" && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j:j+1] == "." {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j:j+1] == s[i:i+1] {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j:j+1] == "*" {
				if p[j-1:j] != s[i:i+1] && p[j-1:j] != "." {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1])
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
