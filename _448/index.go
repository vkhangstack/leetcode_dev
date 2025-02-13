package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	// ways01
	visited := map[int]bool{}
	n := len(nums)
	result := make([]int, 0)
	for _, num := range nums {
		visited[num] = true
	}
	for i := 1; i < n+1; i++ {
		if !visited[i] {
			result = append(result, i)
		}
	}
	return result
}
func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))

}
