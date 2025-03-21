package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)

	rows, cols := n, 2
	paris := make([][]int, rows)

	// Time Complexity: O(n)
	for i := 0; i < n; i++ {
		paris[i] = make([]int, cols)
		paris[i][0] = score[i]
		paris[i][1] = i
	}
	// Time Complexity: O(n.log(n)
	sort.SliceStable(paris, func(i, j int) bool {
		return paris[i][0] > paris[j][0]
	})
	result := make([]string, n)
	// Time Complexity: O(n)
	for i := 0; i < n; i++ {
		index := paris[i][1]

		if i == 0 {
			result[index] = "Gold Medal"
		} else if i == 1 {
			result[index] = "Silver Medal"
		} else if i == 2 {
			result[index] = "Bronze Medal"
		} else {
			result[index] = strconv.Itoa(i + 1)
		}

	}
	return result
}

/**
Total:
Time Complexity: O(n.log(n))
Space Complexity: O(n)
*/

func main() {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
