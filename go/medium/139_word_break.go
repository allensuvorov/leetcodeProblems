* DP
* bottom - up

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true

	for i := n - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if word[0] == s[i] &&
				i+len(word) <= n &&
				word == s[i:i+len(word)] &&
				dp[i+len(word)] {
				dp[i] = true
			}
		}
	}
	return dp[0]
}
