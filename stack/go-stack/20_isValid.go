package main

func isValid(s string) bool {
	openToClose := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	stack := make([]byte, 0)

	for _, runeBracket := range s {
		bracket := byte(runeBracket)
		if bracket == '{' || bracket == '(' || bracket == '[' {
			stack = append(stack, bracket)
		} else if bracket == '}' || bracket == ')' || bracket == ']' {
			if openToClose[stack[len(stack)-1]] == bracket {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}

	return len(stack) == 0
}