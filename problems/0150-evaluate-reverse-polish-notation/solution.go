// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0150

import (
	"strconv"
)

// Problem: Evaluate Reverse Polish Notation
// URL: https://leetcode.com/problems/evaluate-reverse-polish-notation/

func evalRPN(tokens []string) int {
	const maxTokensLen = 10000
	const stackDepth = (maxTokensLen + 1) >> 1

	var operands [stackDepth]int
	var x, y int
	stackSize := 0

	for _, token := range tokens {
		switch token {
		// Operators
		case "+":
			x = operands[stackSize-2]
			y = operands[stackSize-1]
			stackSize--
			operands[stackSize-1] = x + y
		case "-":
			x = operands[stackSize-2]
			y = operands[stackSize-1]
			stackSize--
			operands[stackSize-1] = x - y
		case "*":
			x = operands[stackSize-2]
			y = operands[stackSize-1]
			stackSize--
			operands[stackSize-1] = x * y
		case "/":
			x = operands[stackSize-2]
			y = operands[stackSize-1]
			stackSize--
			operands[stackSize-1] = int(x / y)
		// Operand
		default:
			operands[stackSize], _ = strconv.Atoi(token)
			stackSize++
		}
	}

	return operands[0]
}
