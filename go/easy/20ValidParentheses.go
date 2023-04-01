// need to implement stack using slices
// looks like it's not the only thing needed
// need to review Goland further

func isValid(s string) bool {
	var stack []string
	opn, cls := "({[", ")}]"

	for i := 0; i < len(s); i++ {
		b := string(s[i])
		if strings.Contains(opn, b) {
			stack = append(stack, b) // push
		}
		if strings.Contains(cls, b) {
			if len(stack) == 0 {
				// fmt.Println("stack is empty")
				return false
			} else {
				ind := strings.Index(cls, b)
				n := len(stack) - 1

				if stack[n] == string(opn[ind]) {
					stack = stack[:n] // pop
				} else {
					// fmt.Println("cls with no opn",)
					return false
				}
			}
		}
		// fmt.Println(stack, s[i], b, cls)
	}
	return len(stack) == 0
}