func longestPalindrome(s string) string {
	var res string

	for i := 0; i < len(s); i++ {
		tmp := expandFromMiddle(s, i, i)
		if len(tmp) > len(res) {
			res = tmp
		}
		tmp = expandFromMiddle(s, i, i+1)
		if len(tmp) > len(res) {
			res = tmp
		}
	}

	return res
}

func expandFromMiddle(s string, left int, right int) string {
	var res string
	for l, r := left, right; l >= 0 && r < len(s); l, r = l-1, r+1 {
		if s[l] == s[r] {
			res = s[l : r+1]
		} else {
			break
		}
	}
	return res
}