package parser

import (
	"fmt"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

// EvalRPN take RPN string array as input then parse
// and return int as output
func EvalRPN(tokens []string) (int, error) {
	stk := stack.New()

	for _, token := range tokens {
		if token[0] == '-' && len(token) > 1 || token[0] >= '0' && token[0] <= '9' {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("failed to cast token to int: %v", err)
			}
			stk.Push(num)
			continue
		}

		right := stk.Pop()
		left := stk.Pop()
		if right == nil || left == nil {
			return 0, fmt.Errorf("invailid rpn input")
		}

		if token == "+" {
			stk.Push(left.(int) + right.(int))
		}

		if token == "-" {
			stk.Push(left.(int) - right.(int))
		}

		if token == "*" {
			stk.Push(left.(int) * right.(int))
		}

		if token == "/" {
			stk.Push(left.(int) / right.(int))
		}
	}

	if stk.Len() > 1 {
		return 0, fmt.Errorf("invailid rpn input")
	}

	return stk.Peek().(int), nil
}
