
package main

import "fmt"

type Stack struct {
	items []byte
}

func (s *Stack) IsEmptyStack() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(b byte) {
	s.items = append(s.items, b)
}
func (s *Stack) Pop() byte {
	if !s.IsEmptyStack() {
		element := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return element
	}
	return byte(1)
}
func isValid(s string) bool {
	stk := Stack{}
	for i := 0; i < len(s); i++ {

		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stk.Push(s[i])
		} else if string(s[i]) == ")" || string(s[i]) == "]" || string(s[i]) == "}" {
			result := stk.Pop()
			if result == byte(1) {
				return false
			}
			if (string(result) == "(" && string(s[i]) != ")") || (string(result) == "[" && string(s[i]) != "]") || (string(result) == "{" && string(s[i]) != "}") {
				return false
			}
		}
	}
	if len(stk.items) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Valid Parantheses", isValid("()"))
}
