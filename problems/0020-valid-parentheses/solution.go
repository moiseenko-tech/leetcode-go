// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0020

// Problem: Valid Parentheses
// URL: https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	n := len(s)
	size := 0
	stack := make([]rune, n)

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack[size] = c
			size++
			continue
		}

		if size == 0 {
			return false
		}

		if c == ']' {
			if stack[size-1] != '[' {
				return false
			}
			size--
		}

		if c == '}' {
			if stack[size-1] != '{' {
				return false
			}
			size--
		}

		if c == ')' {
			if stack[size-1] != '(' {
				return false
			}
			size--
		}
	}

	return size == 0
}
