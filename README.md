# LeetCode Solutions in Go

Collection of LeetCode problem solutions implemented in Go.

![LeetCode Progress](https://img.shields.io/badge/LeetCode%20Progress-7%2F3797-yellow)
![Language](https://img.shields.io/badge/Language-Go-00ADD8?logo=go&logoColor=white)

## Structure

```
leetcode-go/
├── .golangci.yml           # golangci-lint configuration
├── LICENSE                 # MIT License
├── Makefile                # Build automation (fmt, lint, test)
├── README.md               # Project documentation
├── go.mod                  # Go module definition
└── problems/
    └── XXXX-problem-name/
        ├── solution.go      # Solution implementation
        └── solution_test.go # Table-driven tests
```

Each problem is organized in its own directory with:
* Solution implementation with problem description, URL, and SPDX license identifier
* Table-driven tests with multiple test cases and parallel execution
* All files include SPDX-License-Identifier headers for license compliance

## Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific problem:
```bash
go test ./problems/0001-two-sum
```

Run tests with verbose output:
```bash
go test -v ./...
```

## Problems Solved

| # | Title | Difficulty | Solution |
|---|-------|------------|----------|
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | 🥉 Easy | [Go](problems/0001-two-sum/solution.go) |
| 4 | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | 🥇 Hard | [Go](problems/0004-median-of-two-sorted-arrays/solution.go) |
| 50 | [Pow(x, n)](https://leetcode.com/problems/powx-n/) | 🥈 Medium | [Go](problems/0050-powx-n/solution.go) | 
| 167 | [Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | 🥈 Medium | [Go](problems/0167-two-sum-ii-input-array-is-sorted/solution.go) |
| 485 | [Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/) | 🥉 Easy | [Go](problems/0485-max-consecutive-ones/solution.go) |
| 961 | [N-Repeated Element in Size 2N Array](https://leetcode.com/problems/n-repeated-element-in-size-2n-array/) | 🥉 Easy | [Go](problems/0961-n-repeated-element-in-size-2n-array/solution.go) |
| 1929 | [Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) | 🥉 Easy | [Go](problems/1929-concatenation-of-array/solution.go) |

## Requirements

- Go 1.25.5 or higher

## License

MIT License - see [LICENSE](LICENSE) file for details.
