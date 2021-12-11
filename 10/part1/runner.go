package aoc

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var count int
	for i := range data {
		stack := make([]byte, 0, len(data[i]))
		var corrupted bool
		for j := range data[i] {
			switch data[i][j] {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack = append(stack, data[i][j])
			case ')':
				n := len(stack) - 1 // Top element
				if stack[n] != '(' {
					corrupted = true
					break
				}
				stack = stack[:n]
			case ']':
				n := len(stack) - 1 // Top element
				if stack[n] != '[' {
					corrupted = true
					break
				}
				stack = stack[:n]
			case '}':
				n := len(stack) - 1 // Top element
				if stack[n] != '{' {
					corrupted = true
					break
				}
				stack = stack[:n]
			case '>':
				n := len(stack) - 1 // Top element
				if stack[n] != '<' {
					corrupted = true
					break
				}
				stack = stack[:n]
			}
			if corrupted {
				switch data[i][j] {
				case ')':
					count += 3
				case ']':
					count += 57
				case '}':
					count += 1197
				case '>':
					count += 25137
				}
				break
			}
		}
	}
	return count
}
