package code_0049

func groupAnagrams(strs []string) [][]string {
	toMap := map[string]int{
		"a": 2,
		"b": 3,
		"c": 5,
		"d": 7,
		"e": 11,
		"f": 13,
		"g": 17,
		"h": 19,
		"i": 23,
		"j": 29,
		"k": 31,
		"l": 37,
		"m": 41,
		"n": 43,
		"o": 47,
		"p": 53,
		"q": 59,
		"r": 61,
		"s": 67,
		"t": 71,
		"u": 73,
		"v": 79,
		"w": 83,
		"x": 89,
		"y": 97,
		"z": 101,
	}

	tMap := make(map[int][]string, 0)
	for i := range strs {
		tValue := 1
		for j := range strs[i] {
			tValue = tValue * toMap[string(strs[i][j])]
		}
		list, ok := tMap[tValue]
		if ok {
			list = append(list, strs[i])
			tMap[tValue] = list
		} else {
			tMap[tValue] = []string{strs[i]}
		}
	}
	result := make([][]string, 0)
	for _, value := range tMap {
		result = append(result, value)
	}
	return result
}
