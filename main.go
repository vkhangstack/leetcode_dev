package main

import (
	"fmt"
	"math"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	indexB := len(b) - 1
	res := make([]byte, len(a))

	var sum, shift byte

	for i := len(a) - 1; i >= 0; i-- {
		sum = shift + a[i]
		if indexB >= 0 {
			sum += b[indexB]
			indexB--
		}

		res[i] = sum%2 + '0'

		shift = sum >> 1 % 2
	}
	if shift > 0 {
		return "1" + string(res)
	}
	return string(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// way1
	//for node := head; node != nil; node = node.Next {
	//	for node.Next != nil && node.Val == node.Next.Val {
	//		node.Next = node.Next.Next
	//	}
	//}
	//return head
	if head == nil {
		return nil
	}

	//way2
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

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

func firstUniqChar(s string) int {
	//O(n+m)
	//visit := make(map[rune]int, len(s))
	//
	//for i := 0; i < len(s); i++ {
	//	visit[rune(s[i])] = visit[rune(s[i])] + 1
	//}
	//
	//for i, v := range s {
	//	if visit[v] == 1 {
	//		return i
	//	}
	//}
	//return -1

	set := [26]int{}
	for i := 0; i < len(s); i++ {
		set[s[i]-'a']++
	}
	fmt.Println(set)

	for i := 0; i < len(s); i++ {
		if set[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	//fmt.Println(addBinary("1", "101"))
	//fmt.Println(isPerfectSquare(16))
	fmt.Println(firstUniqChar("leetcode"))
}
