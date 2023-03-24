package main

import "fmt"

func main() {
	cases := []struct {
		s string
	}{
		{s: "()"},
		{s: "()[]{}"},
		{s: "(]"},
	}

	for _, c := range cases {
		out := isValid(c.s)
		fmt.Printf("Input: s = %s\nOutput: %v\n\n", c.s, out)
	}
}

// leetcode start

func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		default:
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if (top == '(' && v == ')') || (top == '{' && v == '}') || (top == '[' && v == ']') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
