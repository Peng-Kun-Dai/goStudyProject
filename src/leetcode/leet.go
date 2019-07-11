package main

import (
	"fmt"
	"sort"
)

//两数之和
//range
func twoSum(nums []int, target int) []int {
	len := len(nums)
	tar := []int{}
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if target == nums[i]+nums[j] {
				tar = append(tar, i)
				tar = append(tar, j)
			}
		}
	}
	return tar
}

//
func twoSum2(nums []int, target int) []int {
	sort.Ints(nums)

}
func main() {

	nums := []int{2, 7, 11, 15}
	target := 10
	tar := twoSum(nums, target)
	fmt.Println(tar)
}
