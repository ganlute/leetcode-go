package code_5

// Given a string s, return the longest palindromic substring in s
// Input: s = "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.

// dp[0,n] = (s[0] == s[n]) && dp[1, n-1]
// dp[i,i] = true
// dp[i,i+1] = (s[i] == s[i+1])
// dp[i,i+2] = dp[i+1, i+1] && (s[i] == s[i+2])
// dp[i,i+j] = dp[i+1, i+p-1] && (s[i] == s[i+p])

func longestPalindrome(s string) string {
	maxM := 0
	maxN := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			m, n := findLongest(s, i, j)
			if n-m > maxN-maxM {
				maxM = m
				maxN = n
			}
		}
	}
	return s[maxM : maxN+1]
}

func findLongest(s string, i int, j int) (left, right int) {
	mid := i + j
	if (j-i+1)%2 == 1 {
		// 奇数
		left = mid/2 - 1
		right = mid/2 + 1
		return findLongestIndex(s, left, right, i, j)
	} else {
		// 偶数
		left = mid / 2
		right = mid/2 + 1
		return findLongestIndex(s, left, right, i, j)
	}
}
func findLongestIndex(s string, left, right int, i, j int) (int, int) {
	for left >= i && right <= j {
		if s[left] == s[right] {
			left = left - 1
			right = right + 1
		} else {
			break
		}
	}
	left = left + 1
	right = right - 1
	return left, right
}
