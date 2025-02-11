package main

import (
	"fmt"
	_42ValidAnagram "github.com/vkhangstack/leetcode_dev/_242"

	addBinary "github.com/vkhangstack/leetcode_dev/addBinary"
	deleteDuplicates "github.com/vkhangstack/leetcode_dev/deleteDuplicates"
	firstuniqchar "github.com/vkhangstack/leetcode_dev/firstUniqChar"
	isperfectsquare "github.com/vkhangstack/leetcode_dev/isPerfectSquare"
)

func main() {
	fmt.Println(addBinary.Run("1", "101"))
	fmt.Println(deleteDuplicates.Run(&deleteDuplicates.ListNode{Val: 1, Next: &deleteDuplicates.ListNode{Val: 2, Next: nil}}))
	fmt.Println(firstuniqchar.Run("leetcode"))
	fmt.Println(isperfectsquare.Run(16))
	fmt.Println(_42ValidAnagram.Run("rat", "car"))
}
