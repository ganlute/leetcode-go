package code_0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int, 0)
	tMap := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if v, ok := sMap[s[i]]; ok {
			sMap[s[i]] = v + 1
		} else {
			sMap[s[i]] = 1
		}

		if v, ok := tMap[t[i]]; ok {
			tMap[t[i]] = v + 1
		} else {
			tMap[t[i]] = 1
		}
	}

	for key, value := range sMap {
		if tMap[key] != value {
			return false
		}
	}
	return true
}
