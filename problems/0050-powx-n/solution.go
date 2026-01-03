// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Andrei Moiseenko

package p0050

// Problem: Pow(x, n)
// URL: https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n&1 == 0 {
		z := myPow(x, n>>1)
		return z * z
	}

	return x * myPow(x, n-1)
}
