package main

import "fmt"

func findErrorNums(nums []int) []int {
	n := len(nums)
	found := make([]bool, n+1)
	repeatedNum := 0
	runningSum := 0
	desiredSum := (n * (n + 1)) / 2

	for _, num := range nums {
		if found[num] == true {
			if repeatedNum == 0 {
				repeatedNum = num
			}
		} else {
			found[num] = true
			runningSum += num
		}
	}
	missingNum := desiredSum - runningSum
	return []int{repeatedNum, missingNum}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
}
