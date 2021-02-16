func wordBreak(s string, wordDict []string) bool {
	var lookup = make(map[string]bool)
	var dp = make([]bool, len(s)+1)

	for _, word := range wordDict {
		lookup[word] = true
	}

	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := lookup[s[i:j+1]]; dp[i] && ok {
				dp[j+1] = true
			}
		}
	}
	return dp[len(s)]
}