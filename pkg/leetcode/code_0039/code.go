package code_0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if target <= 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	for i := range candidates {
		if target-candidates[i] > 0 {
			list := combinationSum(candidates, target-candidates[i])
			for j := range list {
				r := make([]int, 0)
				r = append(r, candidates[i])
				r = append(r, list[j]...)
				result = append(result, r)
			}
		} else if target-candidates[i] == 0 {
			result = append(result, []int{candidates[i]})
		} else {
			continue
		}
	}
	return deduplicate(result)
}

func deduplicate(arr [][]int) [][]int {
	list := make([][]int, 0)
	for i := range arr {
		if in(list, arr[i]) {
			continue
		} else {
			sort.Ints(arr[i])
			list = append(list, arr[i])
		}
	}
	return list
}
func in(list [][]int, v []int) bool {
	sort.Ints(v)
	for i := range list {
		if len(list[i]) != len(v) {
			continue
		}
		found := true
		for j := 0; j < len(v); j++ {
			if list[i][j] != v[j] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}
