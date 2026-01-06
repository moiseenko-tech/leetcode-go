// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0150

import (
	"testing"
)

var tests = map[string]struct {
	tokens []string
	result int
}{
	"2-1-add-3-mul": {
		tokens: []string{"2", "1", "+", "3", "*"},
		result: 2,
	},
	"4-13-5-div-add": {
		tokens: []string{"4", "13", "5", "/", "+"},
		result: 3,
	},
	"10-6-9-3-add--11-mul-div-mul-17-add-5-add": {
		tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		result: 22,
	},
}

func TestEvalRPN(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := evalRPN(test.tokens)
			expected := test.result
			if got != expected {
				t.Fatalf("evalRPN(tokens=%v) returned: %v; expected: %v",
					test.tokens, got, expected)
			}
		})
	}
}
