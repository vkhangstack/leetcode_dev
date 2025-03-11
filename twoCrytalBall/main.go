package main

import "fmt"

func TwoCrystalBall(breaks []bool) int {
	left := 0
	right := len(breaks) - 1

	for left <= right {
		mid := left + (right-left)/2

		if breaks[mid] {
			break
		} else {
			left = mid + 1
		}
	}
	for i := left; i <= right; i++ {
		if breaks[i] {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(TwoCrystalBall([]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true}))
	fmt.Println(TwoCrystalBall([]bool{true, true, true}))

}
