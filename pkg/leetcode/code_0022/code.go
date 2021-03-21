package code_0022

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	result = append(result, dfs(n, "")...)
	return result
}

func dfs(n int, path string) (result []string) {
	// 提前剪枝
	if !isGoodPath(path) {
		return
	}
	if len(path) == 2*n {
		// 漏网之鱼
		if !isPerfectPath(path) {
			return
		}
		result = append(result, path)
		return
	}
	result = append(result, dfs(n, path+"(")...)
	result = append(result, dfs(n, path+")")...)
	return
}

func isGoodPath(path string) bool {
	st := &stack{}
	st.init()
	for i := range path {
		if string(path[i]) == "(" {
			st.push(string(path[i]))
		} else {
			if "" == st.pop() {
				return false
			}
		}
	}
	return true
}
func isPerfectPath(path string) bool {
	st := &stack{}
	st.init()
	for i := range path {
		if string(path[i]) == "(" {
			st.push(string(path[i]))
		} else {
			if "" == st.pop() {
				return false
			}
		}
	}
	if st.isEmpty() {
		return true
	}
	return false
}

type stack struct {
	arr []string
}

func (this *stack) isEmpty() bool {
	return len(this.arr) == 0
}
func (this *stack) init() {
	this.arr = make([]string, 0)
}
func (this *stack) push(elem string) {
	this.arr = append(this.arr, elem)
}
func (this *stack) pop() (elem string) {
	if len(this.arr) == 0 {
		return ""
	} else {
		elem = this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		return elem
	}
}
