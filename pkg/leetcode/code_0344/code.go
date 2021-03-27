package code_0344

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for j > i {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return
}
