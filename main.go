package main

import "fmt"

// FindDuplicate returns the repeated value in nums, which contains n+1
// integers each in [1, n]. The array is not modified. Values are treated as
// next-pointers (i -> nums[i]), which makes the duplicate the entry point of
// a cycle — found with Floyd's tortoise and hare.
func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	fmt.Println(FindDuplicate([]int{1, 3, 4, 2, 2}))
}
