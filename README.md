# LeetCode Solutions in Go

Collection of LeetCode problem solutions implemented in Go.

![LeetCode Progress](https://img.shields.io/badge/LeetCode%20Progress-10%2F3797-yellow)
![Language](https://img.shields.io/badge/Language-Go-00ADD8?logo=go&logoColor=white)

## Repository structure

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

## Makefile Commands

The project includes a Makefile for common development tasks:

```
make help      # Display all available commands
make all       # Run fmt, lint, and tests (default target)
make fmt       # Format code using golangci-lint
make lint      # Run golangci-lint to check code quality
make lint-fix  # Run golangci-lint with auto-fix
make tests     # Run all tests
```

> [!WARNING]
> `golangci-lint` is required for formatting and linting.
> 
> Install it following the [official documentation](https://golangci-lint.run/docs/).

## Problems Solved

## 🥇 Hard

| # | Title | Difficulty | Solution |
|---|-------|:-----------|:--------:|
| 4 | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | 🥇 Hard | [Go](problems/0004-median-of-two-sorted-arrays/solution.go) |

## 🥈 Medium

|  #  | Title | Difficulty | Solution |
|:---:|-------|:-----------|:--------:|
| 50  | [Pow(x, n)](https://leetcode.com/problems/powx-n/) | 🥈 Medium | [Go](problems/0050-powx-n/solution.go) | 
| 167 | [Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | 🥈 Medium | [Go](problems/0167-two-sum-ii-input-array-is-sorted/solution.go) |

## 🥉 Easy

|  #   | Title | Difficulty | Solution |
|:----:|-------|:-----------|:--------:|
| 1    | [Two Sum](https://leetcode.com/problems/two-sum/) | 🥉 Easy | [Go](problems/0001-two-sum/solution.go) |
| 485  | [Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/) | 🥉 Easy | [Go](problems/0485-max-consecutive-ones/solution.go) |
| 645  | [Set Mismatch](https://leetcode.com/problems/set-mismatch/) | 🥉 Easy | [Go](problems/0645-set-mismatch/solution.go) |
| 961  | [N-Repeated Element in Size 2N Array](https://leetcode.com/problems/n-repeated-element-in-size-2n-array/) | 🥉 Easy | [Go](problems/0961-n-repeated-element-in-size-2n-array/solution.go) |
| 1365 | [How Many Numbers Are Smaller Than the Current Number](https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/) | 🥉 Easy | [Go](problems/1365-how-many-numbers-are-smaller-than-the-current-number/solution.go) |
| 1470 | [Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/) | 🥉 Easy | [Go](problems/1470-shuffle-the-array/) |
| 1929 | [Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) | 🥉 Easy | [Go](problems/1929-concatenation-of-array/solution.go) |

## Requirements

- Go 1.25.5 or higher

## License

MIT License — see [LICENSE](LICENSE) file for details.
