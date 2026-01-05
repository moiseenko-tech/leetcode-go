// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1441

// Problem: Build an Array With Stack Operations
// URL: https://leetcode.com/problems/build-an-array-with-stack-operations/

func buildArray(target []int, n int) []string {
	tgPos, tgLen := 0, len(target)
	tgVal, tgLast := target[tgPos], target[tgLen-1]

	commands := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i == tgVal {
			commands = append(commands, "Push")
			if tgVal == tgLast {
				break
			}
			tgPos++
			tgVal = target[tgPos]
			continue
		}
		commands = append(commands, "Push", "Pop")
	}

	return commands
}
