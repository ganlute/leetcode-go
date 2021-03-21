package code_0020

//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.

const (
	A1 = "("
	A2 = ")"
	B1 = "["
	B2 = "]"
	C1 = "{"
	C2 = "}"
)

var constMap = map[string]string{
	A2: A1,
	B2: B1,
	C2: C1,
}

func isValid(s string) bool {
	st := &stack{}
	st.init()
	for i := range s {
		if _, ok := constMap[string(s[i])]; !ok {
			st.push(string(s[i]))
		} else {
			if st.pop() != constMap[string(s[i])] {
				return false
			}
		}
	}
	if st.isEmpty() {
		return true
	} else {
		return false
	}
}

type stack struct {
	info []string
}

func (this *stack) isEmpty() bool {
	return len(this.info) == 0
}
func (this *stack) init() {
	this.info = make([]string, 0)
}
func (this *stack) push(i string) {
	this.info = append(this.info, i)
}
func (this *stack) pop() string {
	if len(this.info) == 0 {
		return ""
	}
	v := this.info[len(this.info)-1]
	this.info = this.info[:len(this.info)-1]
	return v
}
