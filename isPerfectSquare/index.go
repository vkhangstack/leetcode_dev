package isperfectsquare

import (
	"math"
)

func isPerfectSquare(num int) bool {
	//i := 1
	//
	//for i*i <= num {
	//	if i*i == num {
	//		return true
	//	}
	//	i++
	//}
	//return false

	// O(logN)
	//l, r := 1, num
	//for l <= r {
	//	mid := (l + r) / 2
	//	fmt.Println(mid)
	//	if mid*mid == num {
	//		return true
	//	} else if mid*mid < num {
	//		l = mid + 1
	//	} else {
	//		r = mid - 1
	//	}
	//}
	//return false

	// O(1)
	v := math.Sqrt(float64(num))
	return math.Mod(v, 1) == 0
}

func Run(num int) bool {
	return isPerfectSquare(num)
}
